package request

import "time"

type MakeOrder struct {
	Email string    `json:"email,omitempty"`
	From  time.Time `json:"from,omitempty"`
	To    time.Time `json:"to,omitempty"`
	Room  uint16    `json:"room,omitempty"`
}

func (mo MakeOrder) IsValid() bool {
	return true
}
