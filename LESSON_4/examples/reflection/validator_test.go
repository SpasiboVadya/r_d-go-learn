package reflection

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

var testPerson = Person{
	Username: "john_doe",
	Email:    "john@example.com",
	Age:      30,
}

// BenchmarkValidatorPackage-12    	 1610913	       718.4 ns/op
func BenchmarkValidatorPackage(b *testing.B) {
	validate := validator.New()
	for i := 0; i < b.N; i++ {
		_ = validate.Struct(testPerson)
	}
}

// BenchmarkManualValidation-12    	1000000000	         1.157 ns/op
func BenchmarkManualValidation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testPerson.ValidatePerson()
	}
}
