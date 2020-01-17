// Package order provides the use-case of order module.
package order

import (
	"context"
	"errors"
	entity "gostar/grpc_prevention/order/entity"

	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	// "fmt"
	// "github.com/jinzhu/gorm"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrRecordNotFound record not found
var ErrRecordNotFound = errors.New("product record not found")

// Service is the interface that provides Order methods.
type Service interface {
	// GetByID get store qty from product_route_type
	GetByID(ctx context.Context, id string) (entity.Order, error)
	Sum(ctx context.Context, a int, b int) (int, error)
}

// service implements the Order Service
type service struct {
	repository entity.Repository
	logger     log.Logger
}

// NewService creates and returns a new Order service instance
func NewService(rep entity.Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// GetByID returns an order given by id
func (s *service) GetByID(ctx context.Context, id string) (entity.Order, error) {
	logger := log.With(s.logger, "method", "GetByID")
	order, err := s.repository.GetOrderByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return order, err
}

// Sum returns an sum of 2 integer
func (s *service) Sum(ctx context.Context, a int, b int) (int, error) {
	logger := log.With(s.logger, "method", "Sum")
	c := a + b
	level.Info(logger).Log("info", "sum of a and b")
	return c, nil
}
