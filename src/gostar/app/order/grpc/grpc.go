package grpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	// stdopentracing "github.com/opentracing/opentracing-go"
	// stdzipkin "github.com/openzipkin/zipkin-go"

	orderendpoint "gostar/app/order/endpoint"
	entity "gostar/app/order/entity"
	"gostar/app/order/pb"
	service "gostar/app/order/service"

	"github.com/go-kit/kit/endpoint"

	// "github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/ptypes"
	// "github.com/go-kit/kit/examples/addsvc/pb"
	// "github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	// "github.com/go-kit/kit/examples/addsvc/pkg/addservice"
)

type grpcServer struct {
	sum     grpctransport.Handler
	getbyid grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints orderendpoint.Endpoints, logger log.Logger) pb.OrderServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	// if zipkinTracer != nil {
	// 	// Zipkin GRPC Server Trace can either be instantiated per gRPC method with a
	// 	// provided operation name or a global tracing service can be instantiated
	// 	// without an operation name and fed to each Go kit gRPC server as a
	// 	// ServerOption.
	// 	// In the latter case, the operation name will be the endpoint's grpc method
	// 	// path if used in combination with the Go kit gRPC Interceptor.
	// 	//
	// 	// In this example, we demonstrate a global Zipkin tracing service with
	// 	// Go kit gRPC Interceptor.
	// 	options = append(options, zipkin.GRPCServerTrace(zipkinTracer))
	// }

	return &grpcServer{
		sum: grpctransport.NewServer(
			endpoints.SumEndpoint,
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
			options...,
		),
		getbyid: grpctransport.NewServer(
			endpoints.GetByIDEndpint,
			decodeGRPCGetByIDRequest,
			encodeGRPCGetByIDResponse,
			options...,
		),
	}
}

func (s *grpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumReply, error) {
	_, rep, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SumReply), nil
}

func (s *grpcServer) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDReply, error) {
	_, rep, err := s.getbyid.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetByIDReply), nil
}

// NewGRPCClient returns an AddService backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	// We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.
	// limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	// global client middlewares
	// var options []grpctransport.ClientOption

	// if zipkinTracer != nil {
	// 	// Zipkin GRPC Client Trace can either be instantiated per gRPC method with a
	// 	// provided operation name or a global tracing client can be instantiated
	// 	// without an operation name and fed to each Go kit client as ClientOption.
	// 	// In the latter case, the operation name will be the endpoint's grpc method
	// 	// path.
	// 	//
	// 	// In this example, we demonstrace a global tracing client.
	// 	options = append(options, zipkin.GRPCClientTrace(zipkinTracer))

	// }
	// // Each individual endpoint is an grpc/transport.Client (which implements
	// endpoint.Endpoint) that gets wrapped with various middlewares. If you
	// made your own client library, you'd do this work there, so your server
	// could rely on a consistent set of client behavior.
	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = grpctransport.NewClient(
			conn,
			"pb.Order",
			"Sum",
			encodeGRPCSumRequest,
			decodeGRPCSumResponse,
			pb.SumReply{},
		).Endpoint()
		// sumEndpoint = opentracing.TraceClient(otTracer, "Sum")(sumEndpoint)
		// sumEndpoint = limiter(sumEndpoint)
		// sumEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		// 	Name:    "Sum",
		// 	Timeout: 30 * time.Second,
		// }
		// ))(sumEndpoint)

	}

	var getByIDEndpoint endpoint.Endpoint
	getByIDEndpoint = grpctransport.NewClient(
		conn,
		"pb.Order",
		"GetByID",
		encodeGRPCGetByIDRequest,
		decodeGRPCGetByIDResponse,
		pb.GetByIDReply{},
	).Endpoint()

	// Returning the Endpoints as a service.Service relies on the
	// Endpoints implementing the Service methods. That's just a simple bit
	// of glue code.
	return orderendpoint.Endpoints{
		SumEndpoint:    sumEndpoint,
		GetByIDEndpint: getByIDEndpoint,
	}
}

// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sum request to a user-domain sum request. Primarily useful in a server.
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SumRequest)
	return orderendpoint.SumRequest{A: int(req.A), B: int(req.B)}, nil
}

// decodeGRPCSumResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC sum reply to a user-domain sum response. Primarily useful in a client.
func decodeGRPCSumResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SumReply)
	return orderendpoint.SumResponse{V: int(reply.V), Err: str2err(reply.Err)}, nil
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sum response to a gRPC sum reply. Primarily useful in a server.
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(orderendpoint.SumResponse)
	return &pb.SumReply{V: int64(resp.V), Err: err2str(resp.Err)}, nil
}

// encodeGRPCSumRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain sum request to a gRPC sum request. Primarily useful in a client.
func encodeGRPCSumRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(orderendpoint.SumRequest)
	return &pb.SumRequest{A: int64(req.A), B: int64(req.B)}, nil
}

// decodeGRPCGetByIDRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetByID request to a user-domain GetByID request. Primarily useful in a server.
func decodeGRPCGetByIDRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetByIDRequest)
	return orderendpoint.GetByIDRequest{ID: req.ID}, nil
}

// decodeGRPCGetByIDResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC GetByID reply to a user-domain GetByID response. Primarily useful in a client.
func decodeGRPCGetByIDResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetByIDReply)
	var res entity.Order
	res.ID = reply.Order.ID
	res.PartnerID = reply.Order.PartnerId
	res.State = reply.Order.State
	createDate, _ := ptypes.Timestamp(reply.Order.CreateDate)
	res.CreateDate = createDate
	res.AmountTotal = reply.Order.AmountTotal
	res.WarehouseID = reply.Order.WarehouseId
	//res.OrderItems = reply.Order.OrderItems
	return orderendpoint.GetByIDResponse{Order: res, Err: str2err(reply.Err)}, nil
}

// encodeGRPCGetByIDResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain GetByID response to a gRPC GetByID reply. Primarily useful in a server.
func encodeGRPCGetByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(orderendpoint.GetByIDResponse)
	var res *pb.GetByIDReply_OrderData
	fmt.Println("==========resp")
	fmt.Println(resp)
	fmt.Println(resp.Order)
	fmt.Println(resp.Order.ID)
	res.ID = resp.Order.ID
	res.PartnerId = resp.Order.PartnerID
	res.State = resp.Order.State
	createDate, _ := ptypes.TimestampProto(resp.Order.CreateDate)
	res.CreateDate = createDate
	res.AmountTotal = resp.Order.AmountTotal
	res.WarehouseId = resp.Order.WarehouseID
	return &pb.GetByIDReply{Order: res, Err: err2str(resp.Err)}, nil
}

// encodeGRPCGetByIDRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain GetByID request to a gRPC GetByID request. Primarily useful in a client.
func encodeGRPCGetByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(orderendpoint.GetByIDRequest)
	return &pb.GetByIDRequest{ID: req.ID}, nil
}

// These annoying helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.
// There is special casing to treat empty strings as nil errors.

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
