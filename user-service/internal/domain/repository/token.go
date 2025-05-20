package repository

import (
	"context"
	"time"
)

type TokenRepo interface {
	Put(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
