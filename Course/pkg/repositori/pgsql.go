package repositori

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	mu   sync.Mutex
	pool *pgxpool.Pool
}

func New(commTeg string) (*PGRepo, error) {
	pool, err := pgxpool.Connect(context.Background(), commTeg)
	if err != nil {
		return nil, err
	}

	return &PGRepo{pool: pool, mu: sync.Mutex{}}, err
}
