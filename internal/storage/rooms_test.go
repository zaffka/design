package storage_test

import (
	"errors"
	"testing"
	"time"

	"github.com/zaffka/design/domain"
	"github.com/zaffka/design/internal/storage"
)

func TestRooms_PlaceOrder(t *testing.T) {
	roomStor := storage.NewRooms()

	date1, _ := time.Parse(domain.CustomTimeLayout, "2022-11-10")
	date2, _ := time.Parse(domain.CustomTimeLayout, "2022-11-11")
	date3, _ := time.Parse(domain.CustomTimeLayout, "2022-11-12")
	date4, _ := time.Parse(domain.CustomTimeLayout, "2022-11-13")
	date5, _ := time.Parse(domain.CustomTimeLayout, "2022-11-14")
	date6, _ := time.Parse(domain.CustomTimeLayout, "2022-11-16")

	tests := []struct {
		name  string
		order domain.Order
		err   error
	}{
		{
			name: "order room",
			order: domain.Order{
				UserEmail: "one",
				From:      date2,
				To:        date5,
				Room:      1,
			},
			err: nil,
		},
		{
			name: "same room same time",
			order: domain.Order{
				UserEmail: "one",
				From:      date2,
				To:        date5,
				Room:      1,
			},
			err: storage.ErrRoomOccupied,
		},
		{
			name: "same room From earlier To later",
			order: domain.Order{
				UserEmail: "one",
				From:      date1,
				To:        date6,
				Room:      1,
			},
			err: storage.ErrRoomOccupied,
		},
		{
			name: "same room From later To earlier",
			order: domain.Order{
				UserEmail: "one",
				From:      date3,
				To:        date4,
				Room:      1,
			},
			err: storage.ErrRoomOccupied,
		},
		{
			name: "same room From later To later",
			order: domain.Order{
				UserEmail: "one",
				From:      date5,
				To:        date6,
				Room:      1,
			},
			err: nil,
		},
	}

	tt := tests[0]
	t.Run(tt.name, func(t *testing.T) {
		err := roomStor.PlaceOrder(tt.order)
		if !errors.Is(err, tt.err) {
			t.Fatalf("failed with error: %v", err)
		}
	})

	for _, tt := range tests[1:] {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := roomStor.PlaceOrder(tt.order)
			if !errors.Is(err, tt.err) {
				t.Fatalf("failed with error: %v", err)
			}
		})
	}
}
