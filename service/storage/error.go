package storage

import (
	"github.com/xh3b4sd/anna/service/storage/redis"
)

// IsNotFound combines IsNotFound error matchers of all storage
// implementations. IsNotFound should thus be used for error handling wherever
// spec.Storage is dealt with.
func IsNotFound(err error) bool {
	return redis.IsNotFound(err)
}
