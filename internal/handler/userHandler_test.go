package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brotigen23/gopherMart/internal/config"
	"github.com/brotigen23/gopherMart/internal/dto"
	"github.com/brotigen23/gopherMart/internal/entity"
	"github.com/brotigen23/gopherMart/internal/repository"
	mock_repository "github.com/brotigen23/gopherMart/internal/repository/mocks"
	"github.com/brotigen23/gopherMart/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const target = "localhost:8080/api/user/register"

func TestRegister(t *testing.T) {

	var config = &config.Config{JWTSecretKey: "secret"}
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := mock_repository.NewMockRepository(controller)
	userService := service.NewUserService(mockRepository)

	handler := NewUserHandler(userService, config)

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
					Login:    "user1",
					Password: "any",
				},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusOK,
			},
		}, {
			name: "Test Conflict",
			args: args{
				data: dto.User{
					Login:    "user1",
					Password: "any",
				},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusConflict,
			},
		}, {
			name: "Test Incorrect data",
			args: args{
				data:        dto.User{},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}
	// Test OK [0]
	mockRepository.EXPECT().GetUserByLogin(tests[0].args.data.Login).Return(nil, repository.ErrUserNotFound)
	mockRepository.EXPECT().SaveUser(gomock.Any()).Return(nil, nil)
	// Test Conflict [1]
	mockRepository.EXPECT().GetUserByLogin(tests[1].args.data.Login).Return(nil, nil)

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
		})
	}
}

func TestLogin(t *testing.T) {

	var config = &config.Config{JWTSecretKey: "secret"}
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepository := mock_repository.NewMockRepository(controller)
	userService := service.NewUserService(mockRepository)

	handler := NewUserHandler(userService, config)

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
					Login:    "user1",
					Password: "any",
				},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusOK,
			},
		}, {
			name: "Test User not found",
			args: args{
				data: dto.User{
					Login:    "user2",
					Password: "any",
				},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusConflict,
			},
		}, {
			name: "Test Incorrect data",
			args: args{
				data:        dto.User{},
				contentType: "application/json",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}
	// Test OK [0]
	mockRepository.EXPECT().GetUserByLogin(tests[0].args.data.Login).Return(&entity.User{Login: tests[0].args.data.Login, Password: tests[0].args.data.Password}, nil)
	// Test Conflict [1]
	mockRepository.EXPECT().GetUserByLogin(tests[1].args.data.Login).Return(nil, repository.ErrUserNotFound)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := json.Marshal(tt.args.data)
			assert.NoError(t, err)

			request := httptest.NewRequest(http.MethodPost, target, bytes.NewReader(d))
			request.Header.Set("Content-Type", tt.args.contentType)

			w := httptest.NewRecorder()

			handler.Login(w, request)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}
