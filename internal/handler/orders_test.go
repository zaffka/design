package handler_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zaffka/design/internal/handler"
	"github.com/zaffka/design/internal/storage"
)

func TestOrders(t *testing.T) {
	basePath := "/orders"
	handlr := &handler.Orders{
		RoomManager: storage.NewRooms(),
	}

	type args struct {
		method  string
		path    string
		body    io.Reader
		handler http.Handler
	}

	tests := []struct {
		name              string
		args              args
		expectedCode      int
		expectedInBodyStr string
	}{
		{
			name: "not a get or post request",
			args: args{
				method:  http.MethodPut,
				path:    basePath,
				body:    nil,
				handler: handlr,
			},
			expectedCode:      501,
			expectedInBodyStr: "",
		},
		{
			name: "get request no query",
			args: args{
				method:  http.MethodGet,
				path:    basePath,
				body:    nil,
				handler: handlr,
			},
			expectedCode:      400,
			expectedInBodyStr: "Bad Request",
		},
		{
			name: "get request email query",
			args: args{
				method:  http.MethodGet,
				path:    basePath + "?email=some",
				body:    nil,
				handler: handlr,
			},
			expectedCode:      200,
			expectedInBodyStr: "[]",
		},
		{
			name: "post request",
			args: args{
				method:  http.MethodPost,
				path:    basePath,
				body:    bytes.NewBufferString(`{"email":"some","from":"2022-11-22","to":"2022-11-23","room":1}`),
				handler: handlr,
			},
			expectedCode:      201,
			expectedInBodyStr: "Created",
		},
		{
			name: "get request email query has orders",
			args: args{
				method:  http.MethodGet,
				path:    basePath + "?email=some",
				body:    nil,
				handler: handlr,
			},
			expectedCode:      200,
			expectedInBodyStr: `[{"email":"some","from":`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.args.method, tt.args.path, tt.args.body)
			rec := httptest.NewRecorder()
			tt.args.handler.ServeHTTP(rec, req)

			if rec.Code != tt.expectedCode {
				t.Fatalf("got %d, expected %d", rec.Code, tt.expectedCode)
			}

			body := rec.Body.String()
			if !strings.Contains(body, tt.expectedInBodyStr) {
				t.Fatalf("got body %s, expected body %s", body, tt.expectedInBodyStr)
			}
		})
	}
}
