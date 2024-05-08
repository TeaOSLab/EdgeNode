package ttlcache

type Item[T any] struct {
	Value     T
	expiresAt int64
}
