package entity

import (
	"context"
	"time"
)

// Order represents an order
type Order struct {
	ID          int64       `json:"id"`
	PartnerID   int64       `json:"partner_id"`
	State       string      `json:"state"`
	CreateDate  time.Time   `json:"create_date"`
	AmountTotal float32     `json:"amount_total"`
	WarehouseID int32       `json:"warehouse_id"`
	OrderItems  []OrderItem `json:"order_items,omitempty"`
}

// OrderItem represents items in an order
type OrderItem struct {
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	PriceUnit float32 `json:"price_unit"`
	Quantity  int32   `json:"quantity"`
}

// Repository describes the persistence on order model
type Repository interface {
	GetOrderByID(ctx context.Context, id string) (Order, error)
}
