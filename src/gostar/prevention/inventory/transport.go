package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the inventory service.
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getInventoryHandler := kithttp.NewServer(
		makeGetInventoryEndpoint(bs),
		decodeGetInventoryRequest,
		encodeResponse,
		opts...,
	)
	

	r := mux.NewRouter()

	r.Handle("/inventory/v1/get-inventory/{product_id}", getInventoryHandler).Methods("GET").Queries("product_id","{[0-9]+}")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetInventoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	key := r.FormValue("product_id")
	fmt.Println(key)
	id, ok := vars["product_id"]
	if !ok {
		return nil, errBadRoute
	}
	productId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errBadRoute
	}

	return getInventoryRequest{
		ProductId: productId,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInvalidArgument, RecordNotFound:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}