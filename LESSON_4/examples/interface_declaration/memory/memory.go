package memory

import "fmt"

type MemoryStore struct {
	Data []string
}

func (m MemoryStore) Save(data string) error {
	m.Data = append(m.Data, data)
	fmt.Println("[MemoryStore] Stored in memory:", data)
	return nil
}
