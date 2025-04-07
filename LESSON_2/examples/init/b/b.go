package b

import (
	"fmt"
	_ "lesson02/init/a" // Import package a for side-effects (triggers init)
)

func init() {
	fmt.Println("init from package b")
}

func B() {
	fmt.Println("Function B")
}
