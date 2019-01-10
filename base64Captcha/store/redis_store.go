package store

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

// redisStore An object implementing base64Captcha Store interface.
// when error will be panic.
type redisStore struct {
	client     *redis.Client
	expiration time.Duration
}

// Redis store option.
type Option func(s *redisStore)

// With store expiration option.
func WithExpiration(expiration time.Duration) Option {
	return func(s *redisStore) {
		s.expiration = expiration
	}
}

// With redis Options.
func RedisOption(Option *redis.Options) Option {
	return func(s *redisStore) {
		s.client = redis.NewClient(Option)
	}
}

// Use redis store with Option as base64Captcha Store.
// default is 5 minute duration.
func UseRedisStore(opts ...Option) {
	impl := &redisStore{
		expiration: 5 * time.Minute,
	}

	for _, opt := range opts {
		opt(impl)
	}

	if impl.client == nil {
		impl.client = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default db 0
		})
	}

	base64Captcha.SetCustomStore(impl)
}

// redisStore implementing Set method of Store interface.
// when got error will panic it.
func (s *redisStore) Set(id string, value string) {
	err := s.client.Set(id, value, s.expiration).Err()
	if err != nil {
		panic(err)
	}
}

// redisStore implementing Get method of Store interface.
// when got error will panic it.
func (s *redisStore) Get(id string, clear bool) (value string) {
	val, err := s.client.Get(id).Result()
	if err != nil {
		panic(err)
	}

	if clear {
		err := s.client.Del(id).Err()
		if err != nil {
			panic(err)
		}
	}

	return val
}
