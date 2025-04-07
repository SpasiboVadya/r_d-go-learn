package documentstore

type Collection struct {
	documents map[string]Document
	config    CollectionConfig
}

type CollectionConfig struct {
	PrimaryKey string
}

func NewCollection(cfg CollectionConfig) *Collection {
	return &Collection{
		documents: make(map[string]Document),
		config:    cfg,
	}
}

func (c *Collection) Put(doc Document) {
	keyField, ok := doc.Fields[c.config.PrimaryKey]
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

	c.documents[keyStr] = doc
}

func (c *Collection) Get(key string) (bool, *Document) {
	doc, ok := c.documents[key]
	if !ok {
		return false, nil
	}
	return true, &doc
}

func (c *Collection) Delete(key string) bool {
	_, ok := c.documents[key]
	if ok {
		delete(c.documents, key)
		return true
	}
	return false
}

func (c *Collection) List() []Document {
	docs := make([]Document, 0, len(c.documents))
	for _, doc := range c.documents {
		docs = append(docs, doc)
	}
	return docs
}
