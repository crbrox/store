package dummy

import (
	"fmt"
)

type Store map[string][]byte

func (dummy Store) Put(id string, data []byte) (e error) {
	myData := make([]byte, len(data))
	copy(myData, data)
	dummy[id] = myData
	return nil
}
func (dummy Store) Delete(id string) error {
	delete(dummy, id)
	return nil
}
func (dummy Store) Get(id string) (data []byte, e error) {
	data, ok := dummy[id]
	if !ok {
		e = fmt.Errorf("dummy.Store.Get: %v does not exists", id)
	}
	return data, e
}
func (dummy Store) List() (list []string, e error) {
	for key := range dummy {
		list = append(list, key)
	}
	return list, nil
}
