package redis

import (
	"fmt"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Store struct {
	prefix string
	pool   *redis.Pool
}
type StoreOptions struct {
	Prefix      string
	MaxIdle     int
	MaxActive   int
	Server      string
	IdleTimeout time.Duration
}

func NewStore(options StoreOptions) *Store {

	store := &Store{
		prefix: options.Prefix,
		pool: &redis.Pool{
			MaxIdle:     options.MaxIdle,
			MaxActive:   options.MaxActive,
			IdleTimeout: options.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", options.Server)
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				return nil
				_, err := c.Do("PING")
				return err
			},
		},
	}
	return store
}

func (store *Store) Put(id string, data []byte) error {
	conn := store.pool.Get()
	defer conn.Close()
	_, e := conn.Do("SET", store.prefix+id, data)
	if e != nil {
		return fmt.Errorf("Redis.Put: %v", e)
	}
	return nil
}
func (store *Store) Delete(id string) error {
	conn := store.pool.Get()
	defer conn.Close()
	_, e := conn.Do("DEL", store.prefix+id)
	if e != nil {
		return fmt.Errorf("Redis.Delete: %v", e)
	}
	return nil
}
func (store *Store) Get(id string) (data []byte, e error) {
	conn := store.pool.Get()
	defer conn.Close()
	data, e = redis.Bytes(conn.Do("GET", store.prefix+id))
	return data, e
}
func (store *Store) List() (list []string, e error) {
	conn := store.pool.Get()
	defer conn.Close()
	result, e := redis.Strings(conn.Do("KEYS", store.prefix+"*"))
	for _, elem := range result {
		list = append(list, strings.TrimPrefix(elem, store.prefix))
	}
	return list, nil

}
