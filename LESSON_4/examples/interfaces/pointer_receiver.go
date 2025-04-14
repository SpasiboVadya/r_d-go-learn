package interfaces

import "fmt"

type Closer interface {
	Close() error
}

type File struct{}

func (f File) Close() error {
	fmt.Println("Closed")
	return nil
}

func main() {
	//var c Closer = File{}
	var c Closer = &File{}

	err := c.Close()
	if err != nil {
		fmt.Println(err)
	}
}
