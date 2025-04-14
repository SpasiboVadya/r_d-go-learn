package interface_declaration

import (
	"fmt"
	"lesson04/interface_declaration/memory"
	"math/rand"

	"lesson04/interface_declaration/db"
	"lesson04/interface_declaration/file"
)

// Declared in a business logic or service package
type Store interface {
	Save(data string) error
}

func Process(store Store) {
	err := store.Save("Hello, interface world!")
	if err != nil {
		fmt.Println("Failed to save:", err)
	}
}

func ProcessInMain() {
	var store Store

	switch rand.Intn(3) {
	case 0:
		store = db.DBStore{}
	case 1:
		store = file.FileStore{}
	case 2:
		store = memory.MemoryStore{}
	}

	Process(store)
}
