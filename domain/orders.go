package domain

import "time"

type Order struct {
	UserEmail string    `json:"email,omitempty"`
	From      time.Time `json:"from,omitempty"`
	To        time.Time `json:"to,omitempty"`
	Room      uint64    `json:"room,omitempty"`
}

func (o *Order) Overlaps(newOrder Order) bool {
	return o.To.After(newOrder.From)
}

type OrderList []Order
