package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brotigen23/gopherMart/internal/dto"
	"github.com/brotigen23/gopherMart/internal/service"
	"github.com/stretchr/testify/assert"
)

const target = "localhost:8080/api/user/register"

func TestRegister(t *testing.T) {
	userService := service.NewUserService(nil)
	handler := NewUserHandler(userService, nil)
	type args struct {
		data        dto.User
		contentType string
	}
	type want struct {
		statusCode int
		isCookie   bool
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
				isCookie:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := json.Marshal(tt.args.data)
			assert.NoError(t, err)
			log.Println(string(d))

			request := httptest.NewRequest(http.MethodPost, target, bytes.NewReader(d))
			request.Header.Set("Content-Type", tt.args.contentType)

			w := httptest.NewRecorder()

			handler.Register(w, request)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			if tt.want.isCookie {
				_, err = request.Cookie("token")
				assert.NoError(t, err)
			}
		})
	}
}
