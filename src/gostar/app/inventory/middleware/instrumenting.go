package middleware

import (
	"gostar/app/inventory/service"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	service.Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s service.Service) service.Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) GetInventory(productId int) (quantity float64, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "get-inventory").Add(1)
		s.requestLatency.With("method", "get-inventory").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetInventory(productId)
}
