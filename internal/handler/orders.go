package handler

import (
	"net/http"

	"github.com/zaffka/design/pkg/log"
)

func Orders(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodPost:
		makeOrder(w, r)
	case http.MethodGet:
		getOrders(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("unsupported method used: %s", method)
	}
}

func makeOrder(w http.ResponseWriter, r *http.Request) {

}

func getOrders(w http.ResponseWriter, r *http.Request) {

}
