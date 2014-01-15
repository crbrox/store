package dir

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Store struct{ Path string }

func (store Store) Put(id string, data []byte) (e error) {
	e = ioutil.WriteFile(path.Join(store.Path, id), data, 0660)
	if e != nil {
		return fmt.Errorf("Dir.Put: %v", e)
	}
	return nil
}
func (store Store) Delete(id string) error {
	e := os.Remove(path.Join(store.Path, id))
	if e != nil {
		return fmt.Errorf("Dir.Delete: %v", e)
	}
	return nil
}
func (store Store) Get(id string) (data []byte, e error) {
	data, e = ioutil.ReadFile(path.Join(store.Path, id))
	return data, e
}
func (store Store) List() (list []string, e error) {
	l, e := ioutil.ReadDir(store.Path)
	if e != nil {
		return list, e
	}
	for _, f := range l {
		list = append(list, f.Name())
	}
	return list, nil

}
