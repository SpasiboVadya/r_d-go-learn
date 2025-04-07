//✅ Use reflection when:
// - Reading struct tags for validation, serialization
// - Building libraries or frameworks (ORMs, serializers, validators)
// - You need dynamic type inspection
//❌ Avoid reflection for normal code – it’s slower, unsafe, and harder to read.

package reflection

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyInt int

func BasicReflection() {
	x := MyInt(42)

	val := reflect.ValueOf(x)
	typ := reflect.TypeOf(x)

	fmt.Println("Type:", typ)
	fmt.Println("Kind:", typ.Kind())
	fmt.Println("Value:", val.Int())
}

type User struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"min=18"`
	Email string `json:"email" validate:"required"`
}

func ReadingStructTags() {
	t := reflect.TypeOf(User{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, Tag: %s\n", field.Name, field.Tag)
		fmt.Println("  JSON Tag:", field.Tag.Get("json"))
		fmt.Println("  Validate Tag:", field.Tag.Get("validate"))
	}
}

func validateStruct(s any) []string {
	var errs []string
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("validate")

		// Required check
		if tag == "required" && isEmpty(field) {
			errs = append(errs, fmt.Sprintf("%s is required", t.Field(i).Name))
		}

		// Min check: validate:"min=18"
		if tag != "" && field.Kind() == reflect.Int && hasMinTag(tag) {
			minValue := getMinValue(tag)
			if field.Int() < int64(minValue) {
				errs = append(errs, fmt.Sprintf("%s must be >= %d", t.Field(i).Name, minValue))
			}
		}
	}
	return errs
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int64:
		return v.Int() == 0
	default:
		return false
	}
}

func hasMinTag(tag string) bool {
	return len(tag) > 4 && tag[:4] == "min="
}

func getMinValue(tag string) int {
	val, _ := strconv.Atoi(tag[4:])
	return val
}

func RunValidate() {
	u := User{Name: "", Age: 16, Email: ""}
	errors := validateStruct(u)
	for _, e := range errors {
		fmt.Println("Validation error:", e)
	}
}
