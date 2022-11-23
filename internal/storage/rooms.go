package storage

import (
	"errors"
	"sync"

	"github.com/zaffka/design/domain"
)

var (
	ErrRoomOccupied = errors.New("room already occupied")
)

type roomsData map[uint64]domain.OrderList

type Rooms struct {
	data roomsData
	mu   sync.RWMutex
}

func NewRooms() *Rooms {
	return &Rooms{
		data: make(roomsData),
	}
}

func (r *Rooms) PlaceOrder(newOrder domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderList, ok := r.data[newOrder.Room]
	if !ok {
		r.data[newOrder.Room] = domain.OrderList{
			newOrder,
		}

		return nil
	}

	for _, prevOrder := range orderList {
		if prevOrder.Overlaps(newOrder) {
			return ErrRoomOccupied
		}
	}

	r.data[newOrder.Room] = append(orderList, newOrder)

	return nil
}

func (r *Rooms) GetOrderedByUser(email string) domain.OrderList {
	r.mu.RLock()
	defer r.mu.RUnlock()

	res := domain.OrderList{}
	for _, orders := range r.data {
		for _, order := range orders {
			if order.UserEmail == email {
				res = append(res, order)
			}
		}
	}

	return res
}
