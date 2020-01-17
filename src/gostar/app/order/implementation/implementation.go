package implementation

import (
	"context"
	// "time"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"

	// "github.com/go-kit/kit/log/level"
	// "github.com/gofrs/uuid"
	entity "gostar/app/order/entity"

	"github.com/jinzhu/gorm"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

var RecordNotFound = errors.New("product record not found")

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

// New returns a concrete repository
func New(db *gorm.DB, logger log.Logger) (entity.Repository, error) {
	// return  repository
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "order"),
	}, nil
}

func (rep repository) GetOrderByID(ctx context.Context, id string) (entity.Order, error) {
	order := entity.Order{}
	if id == "" {
		return order, ErrInvalidArgument
	}
	//sqlStatement := `SELECT id, product_id,name,price_unit,product_uom_qty FROM sale_order WHERE client_order_ref = ?`
	sqlStatement := `SELECT id, partner_id,create_date,warehouse_id, state FROM sale_order WHERE client_order_ref = ?`

	err := rep.db.Raw(sqlStatement, id).Scan(&order).Error
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return order, RecordNotFound
	}
	fmt.Println(order)
	fmt.Println(err)
	return order, nil
}
