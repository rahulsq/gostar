package inventory

import (
	"context"
	invserv "gostar/app/inventory/service"
	serv "gostar/app/pkg/service"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestGetInventory(t *testing.T) {
	srv, _ := setup()
	b, err := srv.GetInventory(54549)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing that the date is valid
	if b != 400.0 {
		t.Errorf("inventory is invalid")
	} else {
		t.Logf("inventory is valid")
	}

}

func setup() (srv invserv.Service, ctx context.Context) {
	var db *gorm.DB
	db = serv.GetDBConn()
	return invserv.NewService(db), context.Background()
}
