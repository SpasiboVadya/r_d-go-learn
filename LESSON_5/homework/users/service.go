package users

import (
	"errors"

	"../documentstore"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll *documentstore.Collection
}

func NewService(coll *documentstore.Collection) *Service {
	return &Service{coll: coll}
}

func (s *Service) CreateUser(name string) (*User, error) {
	user := &User{
		ID:   generateID(), // You might want to implement a proper ID generation
		Name: name,
	}

	doc, err := documentstore.MarshalDocument(user)
	if err != nil {
		return nil, err
	}

	s.coll.Put(*doc)
	return user, nil
}

func (s *Service) ListUsers() ([]User, error) {
	docs := s.coll.List()
	users := make([]User, 0, len(docs))

	for _, doc := range docs {
		var user User
		if err := documentstore.UnmarshalDocument(&doc, &user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *Service) GetUser(userID string) (*User, error) {
	found, doc := s.coll.Get(userID)
	if !found {
		return nil, ErrUserNotFound
	}

	var user User
	if err := documentstore.UnmarshalDocument(doc, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) DeleteUser(userID string) error {
	if !s.coll.Delete(userID) {
		return ErrUserNotFound
	}
	return nil
}

// Simple ID generation - you might want to use a better implementation
func generateID() string {
	// This is a simple implementation. In production, you'd want to use a proper ID generation
	return "user_" + string(rune(len(s.coll.List())+1))
}
