package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/zaffka/design/domain"
	"github.com/zaffka/design/internal/handler/request"
	"github.com/zaffka/design/internal/storage"
	"github.com/zaffka/design/pkg/log"
)

type RoomManager interface {
	PlaceOrder(newOrder domain.Order) error
	GetOrderedByUser(email string) domain.OrderList
}

type Orders struct {
	RoomManager RoomManager
}

func (o *Orders) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodPost:
		o.makeOrder(w, r)
	case http.MethodGet:
		o.getOrders(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		log.Errorf("unsupported method used at Orders handler: %s", method)
	}
}

func (o *Orders) makeOrder(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)

	req := &request.MakeOrder{}
	if err := dec.Decode(req); err != nil || !req.IsValid() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	ord := domain.Order{
		UserEmail: req.Email,
		From:      req.From.Time,
		To:        req.To.Time,
		Room:      req.Room,
	}

	err := o.RoomManager.PlaceOrder(ord)
	if errors.Is(err, storage.ErrRoomOccupied) {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)

		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
	status := http.StatusText(http.StatusCreated)

	_, err = w.Write([]byte(status))
	if err != nil {
		log.Errorf("order created but response body writing finished with an error: %s", err)
	}

	log.Infof("order successfully created")
}

func (o *Orders) getOrders(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	orders := o.RoomManager.GetOrderedByUser(userEmail)
	bts, err := json.Marshal(orders)
	if err != nil {
		log.Errorf("error in get orders method: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	_, err = w.Write(bts)
	if err != nil {
		log.Errorf("get orders successfully called but response body writing finished with an error: %s", err)

		return
	}

	log.Infof("get orders successfully called")
}
