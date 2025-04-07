package main

// init() is called automatically before main() — no need to call it yourself.
// It's executed in the order of package dependencies (not lexical or file order).
// Inside a package, multiple init() functions are called in the order they appear in the file(s)
import (
	"fmt"

	"lesson02/homework"

	_ "lesson02/init/c" // still called after b
)

func init() {
	fmt.Println("init from main")
}

func main() {
	fmt.Println(homework.Increment("101"))
}
