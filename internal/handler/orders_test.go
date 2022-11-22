package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zaffka/design/internal/handler"
)

func TestOrders(t *testing.T) {
	anyPath := "/any"

	type args struct {
		method  string
		handler http.HandlerFunc
	}

	tests := []struct {
		name         string
		args         args
		expectedCode int
		expectedBody string
	}{
		{
			name: "not get or post request",
			args: args{
				method:  http.MethodPut,
				handler: handler.Orders,
			},
			expectedCode: 400,
			expectedBody: "",
		},
		{
			name: "get request",
			args: args{
				method:  http.MethodGet,
				handler: handler.Orders,
			},
			expectedCode: 200,
			expectedBody: "",
		},
		{
			name: "post request",
			args: args{
				method:  http.MethodPost,
				handler: handler.Orders,
			},
			expectedCode: 200,
			expectedBody: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(tt.args.method, anyPath, nil)
			rec := httptest.NewRecorder()
			tt.args.handler.ServeHTTP(rec, req)

			if rec.Code != tt.expectedCode {
				t.Fatalf("got %d, expected %d", rec.Code, tt.expectedCode)
			}
		})
	}
}
