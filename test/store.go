package test

import (
	"github.com/crbrox/store"
	"reflect"
	"sort"
	"testing"
)

type Data struct {
	Data []byte
	Id   string
	Ids  []string
}

var d = &Data{
	Data: []byte{1, 2, 3, 4, 6, 'ñ', 'é'},
	Id:   "1234",
	Ids:  []string{"1", "2", "a", "b", "XZ", "YZ"}}

func PutGet(store store.Interface, t *testing.T) {
	PutGetData(store, t, d)
}
func PutGetData(store store.Interface, t *testing.T, d *Data) {
	e := store.Put(d.Id, d.Data)
	if e != nil {
		t.Fatal(e)
	}
	data2, e := store.Get(d.Id)
	if e != nil {
		t.Fatal(e)
	}
	if !reflect.DeepEqual(d.Data, data2) {
		t.Fatalf("expected equal: %v %v", d.Data, data2)
	}
}
func GetInexistent(store store.Interface, t *testing.T) {
	GetInexistentData(store, t, d)
}
func GetInexistentData(store store.Interface, t *testing.T, d *Data) {
	data2, e := store.Get("not exist")
	if e == nil {
		t.Errorf("expected error: %v", data2)
	}
}
func Delete(store store.Interface, t *testing.T) {
	DeleteData(store, t, d)
}
func DeleteData(store store.Interface, t *testing.T, d *Data) {
	e := store.Put(d.Id, d.Data)
	if e != nil {
		t.Fatal(e)
	}
	e = store.Delete(d.Id)
	if e != nil {
		t.Fatal(e)
	}
	data2, e := store.Get(d.Id)
	if e == nil {
		t.Fatalf("expected error: %v %v", data2, e)
	}
}
func List(store store.Interface, t *testing.T) {
	ListData(store, t, d)
}
func ListData(store store.Interface, t *testing.T, d *Data) {
	for _, id := range d.Ids {
		e := store.Put(id, d.Data)
		if e != nil {
			t.Fatal("unexpected error;", e)
		}
	}
	ids2, e := store.List()
	if e != nil {
		t.Fatal(e)
	}
	sort.Strings(d.Ids)
	sort.Strings(ids2)
	if !reflect.DeepEqual(d.Ids, ids2) {
		t.Fatalf("expected equal: %v %v", d.Ids, ids2)
	}
}
func BenchmarkGet(store store.Interface, b *testing.B) {
	BenchmarkGetData(store, b, d)
}
func BenchmarkGetData(store store.Interface, b *testing.B, d *Data) {
	e := store.Put(d.Id, d.Data)
	if e != nil {
		b.Fatal(e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, e := store.Get(d.Id)
		if e != nil {
			b.Fatal(e)
		}
	}
}
func BenchmarkPut(store store.Interface, b *testing.B) {
	BenchmarkPutData(store, b, d)
}
func BenchmarkPutData(store store.Interface, b *testing.B, d *Data) {
	for i := 0; i < b.N; i++ {
		e := store.Put(d.Id, d.Data)
		if e != nil {
			b.Fatal(e)
		}
	}
}
