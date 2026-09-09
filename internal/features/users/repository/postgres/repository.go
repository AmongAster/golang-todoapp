package users_postgres_repository

import core_postgres_pool "github.com/AmongAster/golang-todoapp/internal/core/repository/postgres/pool"

type UsersRepositore struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(
	pool core_postgres_pool.Pool,
) *UsersRepositore {
	return &UsersRepositore{
		pool: pool,
	}
}
