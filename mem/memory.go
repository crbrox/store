package mem

import (
	"fmt"
	"sync"
)

type Store struct {
	data     map[string][]byte
	makeOnce sync.Once
	rwm      sync.RWMutex
}

func (mem *Store) createOnce() {
	mem.makeOnce.Do(func() {
		mem.data = make(map[string][]byte)
	})
}
func (mem *Store) Put(id string, data []byte) (e error) {
	mem.createOnce()
	mem.rwm.Lock()
	defer mem.rwm.Unlock()
	myData := make([]byte, len(data))
	copy(myData, data)
	mem.data[id] = myData
	return nil
}
func (mem *Store) Delete(id string) error {
	mem.createOnce()
	mem.rwm.Lock()
	defer mem.rwm.Unlock()
	delete(mem.data, id)
	return nil
}
func (mem *Store) Get(id string) (data []byte, e error) {
	mem.createOnce()
	mem.rwm.RLock()
	defer mem.rwm.RUnlock()
	data, ok := mem.data[id]
	if !ok {
		e = fmt.Errorf("mem.Store.Get: %v does not exists", id)
	}
	return data, e
}
func (mem *Store) List() (list []string, e error) {
	mem.rwm.Lock()
	defer mem.rwm.Unlock()
	for key := range mem.data {
		list = append(list, key)
	}
	return list, nil
}
