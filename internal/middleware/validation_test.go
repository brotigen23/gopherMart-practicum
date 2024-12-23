package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brotigen23/gopherMart/internal/dto"
	"github.com/stretchr/testify/assert"
)

const target = "localhost:8080/api/user/register"

func TestValidationUser(t *testing.T) {
	type args struct {
		data        dto.User
		contentType string
	}
	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test OK",
			args: args{
				data: dto.User{
					Login:    "123",
					Password: "asd",
				},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "Test Content-type error",
			args: args{
				data:        dto.User{},
				contentType: "plain/text",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Test empty JSON's values",
			args: args{
				data: dto.User{
					Login:    "",
					Password: "",
				},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Test empty JSON",
			args: args{
				data:        dto.User{},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := json.Marshal(tt.args.data)
			assert.NoError(t, err)

			request := httptest.NewRequest(http.MethodPost, target, bytes.NewReader(d))
			request.Header.Set("Content-Type", tt.args.contentType)

			w := httptest.NewRecorder()

			handler := ValidateUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, request)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}
