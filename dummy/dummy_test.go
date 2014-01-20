package dummy

import (
	"github.com/crbrox/store/test"
	"testing"
)

// Generic store tests

func TestPutGet(t *testing.T) {
	var dummy = &Store{}
	test.PutGet(dummy, t)
}
func TestDelete(t *testing.T) {
	var dummy = &Store{}
	test.Delete(dummy, t)
}
func TestList(t *testing.T) {
	var dummy = &Store{}
	test.List(dummy, t)
}
func TestGetInexistent(t *testing.T) {
	var dummy = &Store{}
	test.GetInexistent(dummy, t)
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
	var dummy = &Store{}
	l, e := dummy.List()
	if len(l) != 0 {
		t.Fatalf("l is not empty: %#v", l)
	}
	if e != nil {
		t.Fatalf("unexpected error %#v", e)
	}
}
