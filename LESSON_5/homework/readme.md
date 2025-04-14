1. Потрібно додати функції для маршалінга документів в інші типи.
Тут є два способи:
  1. Використовувати рефлексію (пакет reflect). Ми на лекціях не проходили рефлексію, але за це можна отримати додаткові бали якщо самі розберетесь.
  2. Простіший варінт - використовувати json.Marshal та json.Unmarshal щоб привести вхідний тип до мапи, і з мапи вже будувати наш документ 

```go
func MarshalDocument(input any) (*Document, error) {
  // ..
}

func UnmarshalDocument(doc *Document, output any) error {
  // ..
}
```
Приклади використання:
```go
type MyStruct struct {
  X int
}

func marshalExample() {
  s := &MyStruct{X: 1}
  doc, err := MarshalDocument(s)
  if err != nil {
    fmt.Printf("failed to marshal document: %v\n", err)
    return
  }

  fmt.Printf("marshaled document: %v\n", doc)
}

func unmarshalExample() {
  doc := &Document{
    // ...
  }

  s := &MyStruct{}
  err := UnmarshalDocument(doc, s)
  if err != nil {
    fmt.Printf("failed to unmarshal document: %v\n", err)
    return
  }
 
  fmt.Printf("unmarshaled document: %v\n", s)
}
```
---

2. Потрібно покрити  наші методи щоб повертали відповідні помилки якщо щось пішло не так.
Приклади помилок:
```go
var ErrDocumentNotFound = errors.New("document not found")
var ErrCollectionAlreadyExists = errors.New("collection already exists")
var ErrCollectionNotFound = errors.New("collection not found")
var ErrUnsupportedDocumentField = errors.New("unsupported focument field")
// ...
```
3. Написати users сервіс який використовує наш store 
```go
package users

var (
   ErrUserNotFound = errors.New("user not found")
)

type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

type Service struct {
  coll store.Collection
}

func (s *Service) CreateUser(/* Create user params */) (*User, error) {
  // ...
}

func (s *Service) ListUsers() ([]User, error) {
  // ...
}

func (s *Service) GetUser(userID string) (*User, error) {
  // ...
}

func (s *Service) DeleteUser(userID string) error {
  // ...
}

4. Написати сценарій використання нашого users сервіса у функції main.  
```


NOTE: 

Домашнє завдання виконуємо в директорії lesson_05 
В результаті у нас буде 3 пакета:
main  - lesson_05 
documentstore - lesson_05/document_store 
users - lesson_05/users 

Для зручності можна скопіювати пакет store із попереднього домашнього завдання щоб не писати заново.

---

