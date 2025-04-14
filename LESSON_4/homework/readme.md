# Lesson 4 Методи та інтерфейси

 1.  Потрібно модифікувати пакет documentstore з 3-го заняття відповідно до опису структур та функцій
2. Данні кожної колекції зберігаються в структурі колекції (не в глобальній змінній)

Файл document.go 

package documentstore

type DocumentFieldType string

const (
 DocumentFieldTypeString DocumentFieldType  = "string"
 DocumentFieldTypeNumber DocumentFieldType  = "number"
 DocumentFieldTypeBool   DocumentFieldType  = "bool"
 DocumentFieldTypeArray  DocumentFieldType  = "array"
 DocumentFieldTypeObject DocumentFieldType  = "object"
)

type DocumentField struct {
 Type DocumentFieldType
 // ...
}

type Document struct {
 Fields map[string]DocumentField
}



Файл collection.go 

package documentstore

type Collection struct {
 // ...
}

type CollectionConfig struct {
 PrimaryKey string
}

func (s *Collection) Put(doc Document) {
 // Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`
 // TODO: Implement
}

func (s *Collection) Get(key string) (bool, *Document) {
 // TODO: Implement
}

func (s *Collection) Delete(key string) bool {
 // TODO: Implement
}

func (s *Collection) List() []Document {
 // TODO: Implement
}



Файл store.go 

package documentstore

type Store struct {
 // ...
}

func NewStore() *Store {
 // TODO: Implement
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
 // Створюємо нву колекцію і повертаємо `true` якщо колекція була створена 
 // Якщо ж колекція вже створеня то повертаємо `false` та nil
 // TODO: Implement
}

func (s *Store) GetCollection(name string) (bool, *Collection) {
 // TODO: Implement
}

func (s *Store) DeleteCollection(name string) bool {
 // TODO: Implement
}


