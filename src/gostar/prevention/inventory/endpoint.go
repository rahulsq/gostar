package inventory

import (
	"context"
	"github.com/go-kit/kit/endpoint"

)

type getInventoryRequest struct {
	ProductId int `json:"product_id"`
}

type getInventoryResponse struct {
	Quantity  float64 `json:"quantity"`
	Err error            `json:"error,omitempty"`
}

func (r getInventoryResponse) error() error { return r.Err }

func makeGetInventoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getInventoryRequest)
		quantity, err := s.GetInventory(req.ProductId)
		return getInventoryResponse{Quantity: quantity, Err: err}, nil
	}
}

