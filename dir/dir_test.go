package dir

import (
	"github.com/crbrox/store"
	"github.com/crbrox/store/test"
	"os"
	"testing"
)

const testPath = "testingData"
const notTestPath = "notTestingData"

func createStore(t testing.TB) store.Interface {
	e := os.MkdirAll(testPath, 0770)
	if e != nil {
		t.Fatal(e)
	}
	dir := Store{testPath}
	return dir
}
func delStore(t testing.TB) {
	e := os.RemoveAll(testPath)
	if e != nil {
		t.Fatal(e)
	}
}

//Generic store tests

func TestPutGet(t *testing.T) {
	store := createStore(t)
	defer delStore(t)
	test.PutGet(store, t)
}
func TestDelete(t *testing.T) {
	store := createStore(t)
	defer delStore(t)
	test.Delete(store, t)
}
func TestList(t *testing.T) {
	store := createStore(t)
	defer delStore(t)
	test.List(store, t)
}
func TestGetInexistent(t *testing.T) {
	store := createStore(t)
	defer delStore(t)
	test.GetInexistent(store, t)
}

func BenchmarkGet(b *testing.B) {
	store := createStore(b)
	defer delStore(b)
	test.BenchmarkGet(store, b)
}
func BenchmarkPut(b *testing.B) {
	store := createStore(b)
	defer delStore(b)
	test.BenchmarkGet(store, b)
}

// Specific tests

var aTest = test.Data{
	Data: []byte{},
	Id:   "whatever",
	Ids:  []string{"one", "two", "threee"},
}

func TestPutError(t *testing.T) {
	dir := Store{notTestPath}
	e := dir.Put(aTest.Id, aTest.Data)
	if e == nil {
		t.Fatal("expected error")
	}
}
func TestGetError(t *testing.T) {
	dir := Store{notTestPath}
	_, e := dir.Get(aTest.Id)
	if e == nil {
		t.Fatal("expected error")
	}
}
func TestDeleteError(t *testing.T) {
	dir := Store{notTestPath}
	e := dir.Delete(aTest.Id)
	if e == nil {
		t.Fatal("expected error")
	}
}
func TestListError(t *testing.T) {
	dir := Store{notTestPath}
	_, e := dir.List()
	if e == nil {
		t.Fatal("expected error")
	}
}
