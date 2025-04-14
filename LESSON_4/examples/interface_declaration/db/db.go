package db

import "fmt"

type DBStore struct{}

func (d DBStore) Save(data string) error {
	fmt.Println("[DBStore] Inserting into database:", data)
	return nil
}
