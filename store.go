// store project store.go
package store

type Interface interface {
	Put(id string, data []byte) (e error)
	Delete(id string) error
	Get(id string) ([]byte, error)
	List() ([]string, error)
}
