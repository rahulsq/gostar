package endpoint

import (
	"context"

	"gostar/app/inventory/service"

	"github.com/go-kit/kit/endpoint"
)

type GetInventoryRequest struct {
	ProductId int `json:"product_id"`
}

type GetInventoryResponse struct {
	Quantity float64 `json:"quantity"`
	Err      error   `json:"error,omitempty"`
}

func (r GetInventoryResponse) error() error { return r.Err }

func MakeGetInventoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetInventoryRequest)
		quantity, err := s.GetInventory(req.ProductId)
		return GetInventoryResponse{Quantity: quantity, Err: err}, nil
	}
}
