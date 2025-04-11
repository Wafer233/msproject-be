package cache

import "time"

type Cache interface {
	Put(key string, value string, expire time.Duration) error
	Get(key string) (string, error)
}
