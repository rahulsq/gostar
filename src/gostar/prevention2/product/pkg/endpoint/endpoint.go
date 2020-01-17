package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "gostar/prevention/product/pkg/service"
)

// CreateProductRequest collects the request parameters for the CreateProduct method.
type CreateProductRequest struct {
	Name string `json:"name"`
}

// CreateProductResponse collects the response parameters for the CreateProduct method.
type CreateProductResponse struct {
	Err error `json:"err"`
}

// MakeCreateProductEndpoint returns an endpoint that invokes CreateProduct on the service.
func MakeCreateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		err := s.CreateProduct(ctx, req.Name)
		return CreateProductResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r CreateProductResponse) Failed() error {
	return r.Err
}

// GetProductRequest collects the request parameters for the GetProduct method.
type GetProductRequest struct {
	Name string `json:"name"`
}

// GetProductResponse collects the response parameters for the GetProduct method.
type GetProductResponse struct {
	Id  int   `json:"id"`
	Err error `json:"err"`
}

// MakeGetProductEndpoint returns an endpoint that invokes GetProduct on the service.
func MakeGetProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		id, err := s.GetProduct(ctx, req.Name)
		return GetProductResponse{
			Err: err,
			Id:  id,
		}, nil
	}
}

// Failed implements Failer.
func (r GetProductResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateProduct implements Service. Primarily useful in a client.
func (e Endpoints) CreateProduct(ctx context.Context, name string) (err error) {
	request := CreateProductRequest{Name: name}
	response, err := e.CreateProductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateProductResponse).Err
}

// GetProduct implements Service. Primarily useful in a client.
func (e Endpoints) GetProduct(ctx context.Context, name string) (id int, err error) {
	request := GetProductRequest{Name: name}
	response, err := e.GetProductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetProductResponse).Id, response.(GetProductResponse).Err
}
