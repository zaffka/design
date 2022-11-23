package request_test

import (
	"encoding/json"
	"testing"

	"github.com/zaffka/design/internal/handler/request"
)

func TestMakeOrder_IsValid(t *testing.T) {
	tests := []struct {
		name      string
		orderJSON []byte
		want      bool
	}{
		{
			name:      "ok",
			orderJSON: []byte(`{"email":"some","from":"2022-11-22","to":"2022-11-23","room":1}`),
			want:      true,
		},
		{
			name:      "not ok",
			orderJSON: []byte(`{}`),
			want:      false,
		},
		{
			name:      "wrong time format",
			orderJSON: []byte(`{"email":"some","from":1,"to":1,"room":1}`),
			want:      false,
		},
		{
			name:      "empty email",
			orderJSON: []byte(`{"email":null,"from":"2022-11-22","to":"2022-11-23","room":1}`),
			want:      false,
		},
		{
			name:      "empty room",
			orderJSON: []byte(`{"email":"some","from":"2022-11-22","to":"2022-11-23","room":0}`),
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mo request.MakeOrder
			if err := json.Unmarshal(tt.orderJSON, &mo); err != nil {
				t.Fatal(err)
			}
			if got := mo.IsValid(); got != tt.want {
				t.Errorf("MakeOrder.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
