package errs

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Code    int
}

func (c CustomError) Error() string {
	return fmt.Sprintf("code %d: %s", c.Code, c.Message)
}

// func GetUser() error {
// 	// return &CustomError{}

// 	err := callDB()
// 	if err != nil {
// 		return fmt.Errorf("failed call db: %w", err)
// 	}

// 	err = callMethodb()
// 	if err != nil {
// 		return fmt.Errorf("Error getting user: %w", err)
// 	}

// 	return nil
// 	// return nil
// }

func errorsNew() error {
	return errors.New("this is a simple error")
}

var ErrUserNotFound = errors.New("user not found")

func errorsFormat() error {
	return fmt.Errorf("db err: %v", ErrUserNotFound)
}

func errorsJoin() error {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	return errors.Join(err1, err2)
}

var ErrDbError = errors.New("database error")

func getUser() error {
	if err := callDB(); err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	return nil
}

func MainErrs() {
	err := getUser()
	if err != nil {
		var customErr CustomError
		if errors.As(err, &customErr) {
			fmt.Printf(
				"this is db error: %d, %s\n",
				customErr.Code,
				customErr.Message,
			)
		}
		fmt.Println(err)
	}
}

func ErrorJoins() error {
	errs := make([]error, 0)
	for i := 0; i < 10; i++ {
		errs = append(errs, callDB())
	}
	return errors.Join(errs...)
}

func callDB() *CustomError {
	return nil
}

func CommonPitfallWithCustomError() error {
	err := callDB()
	return err
}

// ce := &CustomError{}
// if errors.As(err, &ce) {
// 	fmt.Println(ce.B)
// 	fmt.Println(ce.Error())
// }

// if errors.Is(err, ErrUserNotFound) {
// 	fmt.Println(err)
// 	return
// }

// if err != nil {
// 	fmt.Println("Something went wrong: %v", err)
// 	return
// }

// fmt.Println("Success")

// if err != nil && err.Error() == "User not found" {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Println("No error")
// }
//}
