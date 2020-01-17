package main

import (
	//  "context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "time"

	"github.com/oklog/oklog/pkg/group"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/jinzhu/gorm"

	"gostar/grpc_prevention/inventory"
	order "gostar/grpc_prevention/order"
	rep "gostar/grpc_prevention/order/implementation"
	serv "gostar/grpc_prevention/pkg/service"

	addpb "gostar/grpc_prevention/order/pb"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

func main() {
	var (
		addr = envString("PORT", defaultPort)
		// rsurl = envString("ROUTINGSERVICE_URL", defaultRoutingServiceURL)

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		// routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")

		// ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var db *gorm.DB
	db = serv.GetDBConn()
	// // Configure some questionable dependencies.
	// var (
	// 	handlingEventFactory = cargo.HandlingEventFactory{
	// 		CargoRepository:    cargos,
	// 		VoyageRepository:   voyages,
	// 		LocationRepository: locations,
	// 	}
	// 	handlingEventHandler = handling.NewEventHandler(
	// 		inspection.NewService(cargos, handlingEvents, nil),
	// 	)
	// )

	// Facilitate testing by adding some cargos.
	//storeTestData(cargos)

	fieldKeys := []string{"method"}

	// var rs routing.Service
	// rs = routing.NewProxyingMiddleware(ctx, *routingServiceURL)(rs)

	var inv inventory.Service
	inv = inventory.NewService(db)
	inv = inventory.NewLoggingService(log.With(logger, "component", "inventory"), inv)
	inv = inventory.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "inventory_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "inventory_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		inv,
	)

	repo, _ := rep.New(db, logger)
	orderService := order.NewService(repo, logger)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/inventory/", inventory.MakeHandler(inv, httpLogger))
	mux.Handle("/order/", order.MakeHandler(orderService, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// grpc server implementation
	grpcAddr := flag.String("grpc-addr", ":8082", "gRPC listen address")
	var g group.Group
	var (
		endpoints  = order.MakeEndpoint(orderService)
		grpcServer = order.NewGRPCServer(endpoints, logger)
	)

	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", *grpcAddr)
			// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
			// the here demonstrated zipkin tracing middleware.
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			addpb.RegisterOrderServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	logger.Log("exit", g.Run())
	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

// func storeTestData(r cargo.Repository) {
// 	test1 := cargo.New("FTL456", cargo.RouteSpecification{
// 		Origin:          location.AUMEL,
// 		Destination:     location.SESTO,
// 		ArrivalDeadline: time.Now().AddDate(0, 0, 7),
// 	})
// 	if err := r.Store(test1); err != nil {
// 		panic(err)
// 	}

// 	test2 := cargo.New("ABC123", cargo.RouteSpecification{
// 		Origin:          location.SESTO,
// 		Destination:     location.CNHcd ..KG,
// 		ArrivalDeadline: time.Now().AddDate(0, 0, 14),
// 	})
// 	if err := r.Store(test2); err != nil {
// 		panic(err)
// 	}
// }
