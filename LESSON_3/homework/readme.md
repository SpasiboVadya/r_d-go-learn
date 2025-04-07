# Lesson 3

1. Потрібно імплементувати пакет document_store відповідно до опису
2. В функції main потрібно описати тестовий сценарій використання стору
3. Данні зберігаємо в глобальній змінній

```go
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

func Put(doc Document) {
 // 1. Перевірити що документ містить поле `key` типу `string`
 // TODO: Implement
}

func Get(key string) (bool, *Document) {
 // Потрібно повернути документ по ключу
 // Якщо документ знайдено, повертаємо `true` та поінтер на документ
 // Інакше повертаємо `false` та `nil`
 // TODO: Implement
}

func Delete(key string) bool { 
 // Видаляємо документа по ключу.
 // Повертаємо `true` якщо ми знайшли і видалили документі
 // Повертаємо `false` якщо документ не знайдено
 // TODO: Implement
}

func List() []Document {
 // Повертаємо список усіх документів
 // TODO: Implement
}
```