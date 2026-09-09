package users_service

import (
	"context"
	"fmt"

	"github.com/AmongAster/golang-todoapp/internal/core/domain"
)

func (s *usersService) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	user, err := s.usersRepositore.GetUser(ctx, id)

	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
