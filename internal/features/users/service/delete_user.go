package users_service

import (
	"context"
	"fmt"
)

func (s *usersService) DeleteUser(
	ctx context.Context,
	id int,
) error {
	if err := s.usersRepositore.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}
