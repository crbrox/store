package redis

import (
	"crbrox/store/test"
	"testing"
	"time"
)

// Generic store tests

var optionsStore = StoreOptions{
	Prefix: "store-test:", MaxIdle: 3, Server: "localhost:6379", IdleTimeout: 240 * time.Second,
}

func TestPutGet(t *testing.T) {
	var store = NewStore(optionsStore)
	test.PutGet(store, t)
}
func TestDelete(t *testing.T) {
	var store = NewStore(optionsStore)
	test.Delete(store, t)
}
func TestList(t *testing.T) {
	var store = NewStore(optionsStore)
	test.List(store, t)
}
func TestGetInexistent(t *testing.T) {
	var store = NewStore(optionsStore)
	test.GetInexistent(store, t)
}

func BenchmarkGet(b *testing.B) {
	var store = NewStore(optionsStore)
	test.BenchmarkGet(store, b)
}

func BenchmarkPut(b *testing.B) {
	var store = NewStore(optionsStore)
	test.BenchmarkPut(store, b)
}
