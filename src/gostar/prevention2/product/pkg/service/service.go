package service

import "context"

// ProductService describes the service.
type ProductService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	CreateProduct(ctx context.Context, name string) (err error)
	GetProduct(ctx context.Context, name string) (id int, err error)
}

type basicProductService struct{}

func (b *basicProductService) CreateProduct(ctx context.Context, name string) (err error) {
	// TODO implement the business logic of CreateProduct
	return err
}
func (b *basicProductService) GetProduct(ctx context.Context, name string) (id int, err error) {
	// TODO implement the business logic of GetProduct
	return id, err
}

// NewBasicProductService returns a naive, stateless implementation of ProductService.
func NewBasicProductService() ProductService {
	return &basicProductService{}
}

// New returns a ProductService with all of the expected middleware wired in.
func New(middleware []Middleware) ProductService {
	var svc ProductService = NewBasicProductService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
