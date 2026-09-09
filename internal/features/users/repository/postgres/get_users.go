package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/AmongAster/golang-todoapp/internal/core/domain"
)

func (r *UsersRepositore) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {
		var userModal UserModel
		if err := rows.Scan(
			&userModal.ID,
			&userModal.Version,
			&userModal.FullName,
			&userModal.PhoneNumber,
		); err != nil {
			return nil, fmt.Errorf("scan users: %w", err)
		}

		userModels = append(userModels, userModal)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	userDomains := userDomainsFromModels(userModels)

	return userDomains, nil
}
