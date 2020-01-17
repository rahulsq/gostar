package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "gostar/prevention/category/pkg/service"
)

// GetByNameRequest collects the request parameters for the GetByName method.
type GetByNameRequest struct {
	S string `json:"s"`
}

// GetByNameResponse collects the response parameters for the GetByName method.
type GetByNameResponse struct {
	Rs  int   `json:"rs"`
	Err error `json:"err"`
}

// MakeGetByNameEndpoint returns an endpoint that invokes GetByName on the service.
func MakeGetByNameEndpoint(s service.CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByNameRequest)
		rs, err := s.GetByName(ctx, req.S)
		return GetByNameResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByNameResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetByName implements Service. Primarily useful in a client.
func (e Endpoints) GetByName(ctx context.Context, s string) (rs int, err error) {
	request := GetByNameRequest{S: s}
	response, err := e.GetByNameEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByNameResponse).Rs, response.(GetByNameResponse).Err
}
