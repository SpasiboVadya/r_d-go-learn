package main

import (
	"fmt"

	"./documentstore"
)

func main() {
	// Create a new store
	store := documentstore.NewStore()

	// Create a collection with "id" as primary key
	cfg := &documentstore.CollectionConfig{
		PrimaryKey: "id",
	}
	created, usersCollection := store.CreateCollection("users", cfg)
	if created {
		fmt.Println("Created 'users' collection")
	}

	// Create a document
	userDoc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "user1",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "John Doe",
			},
			"age": {
				Type:  documentstore.DocumentFieldTypeNumber,
				Value: 30,
			},
			"isActive": {
				Type:  documentstore.DocumentFieldTypeBool,
				Value: true,
			},
		},
	}

	// Put the document in the collection
	usersCollection.Put(userDoc)

	// Get the document
	found, doc := usersCollection.Get("user1")
	if found {
		fmt.Printf("Found user document: %+v\n", doc.Fields)
	}

	// List all documents in the collection
	allDocs := usersCollection.List()
	fmt.Println("\nAll documents in users collection:")
	for _, doc := range allDocs {
		fmt.Printf("%+v\n", doc.Fields)
	}

	// Try to create the same collection again
	created, _ = store.CreateCollection("users", cfg)
	if !created {
		fmt.Println("\nFailed to create 'users' collection again (as expected)")
	}

	// Delete the collection
	deleted := store.DeleteCollection("users")
	if deleted {
		fmt.Println("Successfully deleted 'users' collection")
	}

	// Try to get the deleted collection
	found, _ = store.GetCollection("users")
	if !found {
		fmt.Println("Confirmed 'users' collection is deleted")
	}
}
