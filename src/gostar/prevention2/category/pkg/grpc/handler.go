package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "gostar/prevention/category/pkg/endpoint"
	pb "gostar/prevention/category/pkg/grpc/pb"
)

// makeGetByNameHandler creates the handler logic
func makeGetByNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByNameEndpoint, decodeGetByNameRequest, encodeGetByNameResponse, options...)
}

// decodeGetByNameResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByName request.
// TODO implement the decoder
func decodeGetByNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Category' Decoder is not impelemented")
}

// encodeGetByNameResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Category' Encoder is not impelemented")
}
func (g *grpcServer) GetByName(ctx context1.Context, req *pb.GetByNameRequest) (*pb.GetByNameReply, error) {
	_, rep, err := g.getByName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetByNameReply), nil
}
