package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	// "strconv"
	// "fmt"

	orderendpoint "gostar/app/order/endpoint"
	service "gostar/app/order/service"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the inventory service.
func MakeHandler(order service.Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getByIDHandler := kithttp.NewServer(
		orderendpoint.MakeGetByIDEndpoint(order),
		decodeGetByIDRequest,
		encodeResponse,
		opts...,
	)

	sumHandler := kithttp.NewServer(
		orderendpoint.MakeSumEndpoint(order),
		decodeSumRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/order/v1/{id}", getByIDHandler).Methods("GET")
	r.Handle("/order/v1/sum/{a}/{b}", sumHandler).Methods("GET")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return orderendpoint.GetByIDRequest{ID: id}, nil
}

func decodeSumRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	a, ok := vars["a"]
	if !ok {
		return nil, errBadRoute
	}
	a1, _ := strconv.Atoi(a)
	b, ok := vars["b"]
	if !ok {
		return nil, errBadRoute
	}
	b1, _ := strconv.Atoi(b)
	return orderendpoint.SumRequest{A: a1, B: b1}, nil
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
	case service.ErrInvalidArgument, service.ErrRecordNotFound:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
