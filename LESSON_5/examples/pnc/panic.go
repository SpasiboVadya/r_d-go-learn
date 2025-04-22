package pnc

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("User not found")

type User struct {
	Name string
}

func GetUser() (*User, error) {
	// ...
	return nil, ErrUserNotFound
}

func MustGetUser() {
	panic("panics (╯°□°)╯︵ ┻━┻")
}

func MainPnc() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		} else {
			fmt.Println("No panic ┬─┬ノ(º_ºノ)")
		}
	}()

	//go MustGetUser()

	//time.Sleep(5 * time.Second)

	// err := MustGetUser()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
