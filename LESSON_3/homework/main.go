package main

import (
	"fmt"
	"homework/documentstore"
)

func main() {
	doc1 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "doc1",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Example Document 1",
			},
		},
	}

	documentstore.Put(doc1)

	found, d := documentstore.Get("doc1")
	if found {
		fmt.Printf("Found document with key 'doc1': %+v\n", d.Fields)
	} else {
		fmt.Println("Document with key 'doc1' not found")
	}

	doc2 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "doc2",
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

	documentstore.Put(doc2)

	deleted := documentstore.Delete("doc1")
	if deleted {
		fmt.Println("Document with key 'doc1' was deleted successfully")
	} else {
		fmt.Println("Failed to delete document with key 'doc1'")
	}

	// Спробуємо ще раз отримати doc1, щоб перевірити, що він видалений
	found, _ = documentstore.Get("doc1")
	if !found {
		fmt.Println("Document with key 'doc1' is confirmed deleted")
	}

	// Отримаємо та виведемо список усіх документів у store
	allDocs := documentstore.List()
	fmt.Printf("List of all documents in store:\n")
	for _, doc := range allDocs {
		fmt.Printf("%+v\n", doc.Fields)
	}
}
