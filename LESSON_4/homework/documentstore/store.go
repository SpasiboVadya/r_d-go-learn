package documentstore

type Store struct {
	collections map[string]*Collection
}

func NewStore() *Store {
	return &Store{
		collections: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	if _, exists := s.collections[name]; exists {
		return false, nil
	}

	collection := NewCollection(*cfg)
	s.collections[name] = collection
	return true, collection
}

func (s *Store) GetCollection(name string) (bool, *Collection) {
	collection, exists := s.collections[name]
	if !exists {
		return false, nil
	}
	return true, collection
}

func (s *Store) DeleteCollection(name string) bool {
	if _, exists := s.collections[name]; exists {
		delete(s.collections, name)
		return true
	}
	return false
}
