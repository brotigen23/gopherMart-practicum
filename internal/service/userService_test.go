package service

import (
	"testing"

	"github.com/brotigen23/gopherMart/internal/entity"
	"github.com/brotigen23/gopherMart/internal/repository"
)

func TestSaveUser(t *testing.T) {
	type args struct {
		user entity.User
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test save normal",
			args: args{
				user: entity.User{},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "Test user already exists",
			args: args{
				user: entity.User{},
			},
			want: want{
				err: repository.ErrUserExists,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
