// Package inventory provides the use-case of inventory module.
package inventory

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

var RecordNotFound = errors.New("product record not found")

// Service is the interface that provides inventory methods.
type Service interface {
	// GetInventory get store qty from product_route_type
	GetInventory(productId int) (float64, error)

}

type service struct {
	db         *gorm.DB
}

func (s *service) GetInventory(productId int) (float64, error) {
	if productId <= 0 {
		return 0.0, ErrInvalidArgument
	}
	sqlStatement := `SELECT store_qty FROM product_route_type WHERE id = ?`
	type Quantity struct {
		StoreQty float64
	}
	var quantity Quantity
	err := s.db.Raw(sqlStatement, productId).Scan(&quantity).Error
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return 0.0, RecordNotFound
	}
	fmt.Println(quantity)
	fmt.Println(err)
	return quantity.StoreQty, nil
	
}

// NewService creates a inventory service with necessary dependencies.
func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

