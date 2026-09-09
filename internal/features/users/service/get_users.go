package users_service

import (
	"context"
	"fmt"

	"github.com/AmongAster/golang-todoapp/internal/core/domain"
	core_errors "github.com/AmongAster/golang-todoapp/internal/core/errors"
)

func (s *usersService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrinvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrinvalidArgument,
		)
	}

	users, err := s.usersRepositore.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users from repository: %w", err)
	}

	return users, nil
}
