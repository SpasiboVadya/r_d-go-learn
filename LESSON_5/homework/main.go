package main

import (
	"fmt"

	"./documentstore"
	"./users"
)

func main() {
	// Create a new store
	store := documentstore.NewStore()

	// Create a collection for users
	cfg := &documentstore.CollectionConfig{
		PrimaryKey: "id",
	}
	created, usersCollection := store.CreateCollection("users", cfg)
	if !created {
		fmt.Println("Failed to create users collection")
		return
	}

	// Create users service
	userService := users.NewService(usersCollection)

	// Create some users
	user1, err := userService.CreateUser("John Doe")
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		return
	}
	fmt.Printf("Created user: %+v\n", user1)

	user2, err := userService.CreateUser("Jane Smith")
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		return
	}
	fmt.Printf("Created user: %+v\n", user2)

	// List all users
	allUsers, err := userService.ListUsers()
	if err != nil {
		fmt.Printf("Failed to list users: %v\n", err)
		return
	}
	fmt.Println("\nAll users:")
	for _, user := range allUsers {
		fmt.Printf("%+v\n", user)
	}

	// Get a specific user
	foundUser, err := userService.GetUser(user1.ID)
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		return
	}
	fmt.Printf("\nFound user: %+v\n", foundUser)

	// Delete a user
	err = userService.DeleteUser(user1.ID)
	if err != nil {
		fmt.Printf("Failed to delete user: %v\n", err)
		return
	}
	fmt.Println("\nDeleted user:", user1.ID)

	// Try to get the deleted user
	_, err = userService.GetUser(user1.ID)
	if err == users.ErrUserNotFound {
		fmt.Println("Confirmed user is deleted")
	}
}
