package documentstore

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

var store = make(map[string]Document)

func Put(doc Document) {
	keyField, ok := doc.Fields["key"]
	if !ok {
		return
	}

	if keyField.Type != DocumentFieldTypeString {
		return
	}

	keyStr, ok := keyField.Value.(string)
	if !ok {
		return
	}

	store[keyStr] = doc
}

func Get(key string) (bool, *Document) {
	doc, ok := store[key]
	if !ok {
		return false, nil
	}
	return true, &doc
}

func Delete(key string) bool {
	_, ok := store[key]
	if ok {
		delete(store, key)
		return true
	}
	return false
}

func List() []Document {
	docs := make([]Document, 0, len(store))
	for _, doc := range store {
		docs = append(docs, doc)
	}
	return docs
}
