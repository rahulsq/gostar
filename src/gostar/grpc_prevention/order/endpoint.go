package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	entity "gostar/grpc_prevention/order/entity"
)

// GetByIDRequest holds the request parameters for the GetByID method.
type GetByIDRequest struct {
	ID string
}

// GetByIDResponse holds the response values for the GetByID method.
type GetByIDResponse struct {
	Order entity.Order `json:"order"`
	Err   error        `json:"error,omitempty"`
}

// SumRequest holds the request parameters for the Sum method.
type SumRequest struct {
	A, B int
}

// SumResponse holds the response values for the Sum method.
type SumResponse struct {
	V   int   `json:"sum"`
	Err error `json:"error,omitempty"`
}

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	GetByIDEndpint endpoint.Endpoint
	SumEndpoint    endpoint.Endpoint
}

// MakeEndpoint initializes all Go kit endpoints for the Order service.
func MakeEndpoint(s Service) Endpoints {
	return Endpoints{
		GetByIDEndpint: makeGetByIDEndpoint(s),
		SumEndpoint:    makeSumEndpoint(s),
	}
}

func makeGetByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		orderRes, err := s.GetByID(ctx, req.ID)
		return GetByIDResponse{Order: orderRes, Err: err}, nil
	}
}

func makeSumEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		sum, err := s.Sum(ctx, req.A, req.B)
		return SumResponse{V: sum, Err: err}, nil
	}
}

// Sum implements the service interface, so Set may be used as a service.
// This is primarily useful in the context of a client library.
func (s Endpoints) Sum(ctx context.Context, a, b int) (int, error) {
	resp, err := s.SumEndpoint(ctx, SumRequest{A: a, B: b})
	if err != nil {
		return 0, err
	}
	response := resp.(SumResponse)
	return response.V, response.Err
}

// GetByID implements the service interface, so Endpoints may be used as a service.
// This is primarily useful in the context of a client library.
func (s Endpoints) GetByID(ctx context.Context, a string) (entity.Order, error) {
	var odr entity.Order
	resp, err := s.GetByIDEndpint(ctx, GetByIDRequest{ID: a})
	if err != nil {
		return odr, err
	}
	response := resp.(GetByIDResponse)
	return response.Order, response.Err
}
