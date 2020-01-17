package service

import "context"

// CategoryService describes the service.
type CategoryService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetByName(ctx context.Context, s string) (rs int, err error)
}

type basicCategoryService struct{}

func (b *basicCategoryService) GetByName(ctx context.Context, s string) (rs int, err error) {
	// TODO implement the business logic of GetByName
	return rs, err
}

// NewBasicCategoryService returns a naive, stateless implementation of CategoryService.
func NewBasicCategoryService() CategoryService {
	return &basicCategoryService{}
}

// New returns a CategoryService with all of the expected middleware wired in.
func New(middleware []Middleware) CategoryService {
	var svc CategoryService = NewBasicCategoryService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
