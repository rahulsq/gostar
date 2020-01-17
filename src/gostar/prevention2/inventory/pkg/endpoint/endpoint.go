package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "gostar/prevention/inventory/pkg/service"
	"fmt"
)

// ListExceptionRequest collects the request parameters for the ListException method.
type ListExceptionRequest struct {
	S string `json:"s"`
}

// ListExceptionResponse collects the response parameters for the ListException method.
type ListExceptionResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeListExceptionEndpoint returns an endpoint that invokes ListException on the service.
func MakeListExceptionEndpoint(s service.InventoryService) endpoint.Endpoint {
	fmt.Println("endpoint 1 invoke")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListExceptionRequest)
		rs, err := s.ListException(ctx, req.S)
		return ListExceptionResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ListExceptionResponse) Failed() error {
	fmt.Println("endpoint 2 failed")
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// ListException implements Service. Primarily useful in a client.
func (e Endpoints) ListException(ctx context.Context, s string) (rs interface{}, err error) {
	fmt.Println("endpint 3 client")
	request := ListExceptionRequest{S: s}
	response, err := e.ListExceptionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListExceptionResponse).Rs, response.(ListExceptionResponse).Err
}
