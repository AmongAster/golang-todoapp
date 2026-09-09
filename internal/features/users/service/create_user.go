package users_service

import (
	"context"
	"fmt"

	"github.com/AmongAster/golang-todoapp/internal/core/domain"
)

func (s *usersService) CreateUsers(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validater user domain: %w", err)
	}

	user, err := s.usersRepositore.CreateUsers(ctx, user)

	if err != nil {
		return domain.User{}, fmt.Errorf("cteate user: %w", err)
	}

	return user, nil
}
