package mem

import (
	"github.com/crbrox/store/test"
	"testing"
)

// Generic store tests

func TestPutGet(t *testing.T) {
	var mem = NewStore()
	test.PutGet(mem, t)
}
func TestDelete(t *testing.T) {
	var mem = NewStore()
	test.Delete(mem, t)
}
func TestList(t *testing.T) {
	var mem = NewStore()
	test.List(mem, t)
}
func TestGetInexistent(t *testing.T) {
	var mem = NewStore()
	test.GetInexistent(mem, t)
}

func BenchmarkGet(b *testing.B) {
	var store = NewStore()
	test.BenchmarkGet(store, b)
}

func BenchmarkPut(b *testing.B) {
	var store = NewStore()
	test.BenchmarkPut(store, b)
}

//Specific tests

func TestListEmpty(t *testing.T) {
	var mem = &Store{}
	l, e := mem.List()
	if len(l) != 0 {
		t.Fatalf("l is not empty: %#v", l)
	}
	if e != nil {
		t.Fatalf("unexpected error %#v", e)
	}
}
