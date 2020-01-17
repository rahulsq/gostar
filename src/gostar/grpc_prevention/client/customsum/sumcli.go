package main

import (
	"context"
	"flag"
	"fmt"

	"os"
	"text/tabwriter"
	"time"

	"github.com/go-kit/kit/log"
	lightstep "github.com/lightstep/lightstep-tracer-go"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	order "gostar/grpc_prevention/order"
)

func main() {
	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an addsvc. This presumption is reflected in
	// the addcli binary and the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings. For an
	// example service with a client built on top of a service discovery system,
	// see profilesvc.
	fs := flag.NewFlagSet("addcli", flag.ExitOnError)

	var (
		zipkinURL      = fs.String("zipkin-url", "", "Enable Zipkin tracing via HTTP reporter URL e.g. http://localhost:9411/api/v2/spans")
		zipkinBridge   = fs.Bool("zipkin-ot-bridge", false, "Use Zipkin OpenTracing bridge instead of native implementation")
		lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
		appdashAddr    = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")
	)
	var (
		grpcAddr = fs.String("grpc-addr", "localhost:8082", "GRPC address of addsvc")
	)

	// This is a demonstration of the native Zipkin tracing client. If using
	// Zipkin this is the more idiomatic client over OpenTracing.
	var zipkinTracer *zipkin.Tracer
	{
		if *zipkinURL != "" {
			var (
				err         error
				hostPort    = "" // if host:port is unknown we can keep this empty
				serviceName = "addsvc-cli"
				reporter    = zipkinhttp.NewReporter(*zipkinURL)
			)
			defer reporter.Close()
			zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
			zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP))
			if err != nil {
				fmt.Fprintf(os.Stderr, "unable to create zipkin tracer: %s\n", err.Error())
				os.Exit(1)
			}
		}
	}

	// This is a demonstration client, which supports multiple tracers.
	// Your clients will probably just use one tracer.
	var otTracer stdopentracing.Tracer
	{
		if *zipkinBridge && zipkinTracer != nil {
			otTracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil // do not instrument with both native and ot bridge
		} else if *lightstepToken != "" {
			otTracer = lightstep.NewTracer(lightstep.Options{
				AccessToken: *lightstepToken,
			})
			defer lightstep.FlushLightStepTracer(otTracer)
		} else if *appdashAddr != "" {
			otTracer = appdashot.NewTracer(appdash.NewRemoteCollector(*appdashAddr))
		} else {
			otTracer = stdopentracing.GlobalTracer() // no-op
		}
	}

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	var svc order.Service
	svc = order.NewGRPCClient(conn, log.NewNopLogger())

	a := 4
	b := 5
	fmt.Println(a, b)
	v, err := svc.Sum(context.Background(), int(a), int(b))
	fmt.Println(err)
	fmt.Println(v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%d + %d = %d\n", a, b, v)

	// order details
	id := "111111172"
	fmt.Println(id)
	v1, err1 := svc.GetByID(context.Background(), id)
	fmt.Println(err1)
	fmt.Println(v1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s \n", id)
	fmt.Println(v1)

}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
