package request

import "github.com/zaffka/design/domain"

type MakeOrder struct {
	Email string            `json:"email,omitempty"`
	From  domain.CustomTime `json:"from,omitempty"`
	To    domain.CustomTime `json:"to,omitempty"`
	Room  uint64            `json:"room,omitempty"`
}

func (mo MakeOrder) IsValid() bool {
	return mo.Email != "" && !mo.From.IsZero() && !mo.To.IsZero() && mo.Room > 0
}
