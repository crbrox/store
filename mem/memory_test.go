package mem

import (
	"crbrox/store/test"
	"testing"
)

// Generic store tests

func TestPutGet(t *testing.T) {
	var mem = &Store{}
	test.PutGet(mem, t)
}
func TestDelete(t *testing.T) {
	var mem = &Store{}
	test.Delete(mem, t)
}
func TestList(t *testing.T) {
	var mem = &Store{}
	test.List(mem, t)
}
func TestGetInexistent(t *testing.T) {
	var mem = &Store{}
	test.GetInexistent(mem, t)
}

func BenchmarkGet(b *testing.B) {
	var store = &Store{}
	test.BenchmarkGet(store, b)
}

func BenchmarkPut(b *testing.B) {
	var store = &Store{}
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
