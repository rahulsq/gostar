package middleware

import (
	"gostar/app/inventory/service"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	service.Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s service.Service) service.Service {
	return &loggingService{logger, s}
}

func (s *loggingService) GetInventory(productId int) (quantity float64, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get-inventory",
			"product_id", productId,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetInventory(productId)
}
