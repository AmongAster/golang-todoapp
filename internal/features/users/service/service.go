package users_service

import (
	"context"

	"github.com/AmongAster/golang-todoapp/internal/core/domain"
)

type usersService struct {
	usersRepositore UsersRepositore
}

type UsersRepositore interface {
	CreateUsers(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error
}

func NewUsersService(
	usersRepositore UsersRepositore,

) *usersService {
	return &usersService{
		usersRepositore: usersRepositore,
	}
}
