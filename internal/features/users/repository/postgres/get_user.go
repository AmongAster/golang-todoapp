package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/AmongAster/golang-todoapp/internal/core/domain"
	core_errors "github.com/AmongAster/golang-todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepositore) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var userModal UserModel

	err := row.Scan(
		&userModal.ID,
		&userModal.Version,
		&userModal.FullName,
		&userModal.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user whith id='%d': %w", id, core_errors.ErrNotFound)
		}

		return domain.User{}, fmt.Errorf("scan error: %w ", err)
	}

	userDomain := domain.NewUser(
		userModal.ID,
		userModal.Version,
		userModal.FullName,
		userModal.PhoneNumber,
	)

	return userDomain, nil
}
