package file

import "fmt"

type FileStore struct{}

func (f FileStore) Save(data string) error {
	fmt.Println("[FileStore] Writing to file:", data)
	return nil
}
