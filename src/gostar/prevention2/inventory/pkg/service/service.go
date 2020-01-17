package service

import "context"
import "fmt"

// InventoryService describes the service.
type InventoryService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	ListException(ctx context.Context, s string) (rs interface{}, err error)
}

type basicInventoryService struct{}

func (b *basicInventoryService) ListException(ctx context.Context, s string) (rs interface{}, err error) {
	// TODO implement the business logic of ListException
	rs1 := make(map[string]string, 0)
	rs1["sdds"] = "dsds"
	rs = rs1
	fmt.Println("statttttttttttt")
	return rs, err
}

// NewBasicInventoryService returns a naive, stateless implementation of InventoryService.
func NewBasicInventoryService() InventoryService {
	return &basicInventoryService{}
}

// New returns a InventoryService with all of the expected middleware wired in.
func New(middleware []Middleware) InventoryService {
	var svc InventoryService = NewBasicInventoryService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
