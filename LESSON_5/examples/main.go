package main

import "fmt"

func handleRequest() (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recovered from panic: %v\n", e)
			err = fmt.Errorf("recovered from panic: %v", e)
		}
	}()

	// handle request ...

	return nil
}

var allowedStatuses = map[string]struct{}{
	"running":   {},
	"succeeded": {},
	"failed":    {},
}

func updateStatus(status string) {
	if _, ok := allowedStatuses[status]; !ok {
		panic("invalid status")
	}
}

func main() {
	handleRequest()
	updateStatus("some")
}

//pnc.MainPnc()
//fmt.Println(errs.ErrorJoins())
//errs.MainErrs()
