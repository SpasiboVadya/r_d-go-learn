package interfaces

import (
	"fmt"
	"math/rand"
)

type Animal interface {
	Speak()
}

type Cat struct{}

func (c *Cat) Speak() {
	fmt.Println("Meow")
}

func FindCat() Animal {
	if rand.Intn(2) == 0 {
		return nil
	} else {
		return &Cat{}
	}
}

func NilInterfacePitfall() {
	var animal Animal
	animal = FindCat()
	if animal == nil {
		fmt.Println("Нема кота")
	} else {
		fmt.Println("Погладь кота!")
	}
}
