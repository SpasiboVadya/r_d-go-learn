package recover

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!") // Causes a panic
	fmt.Println("This line won't be reached")
}
