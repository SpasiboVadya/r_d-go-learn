package main

/*
#cgo CFLAGS: -g -Wall
#include <stdlib.h>
int get_string(char **out);
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	var cstr *C.char

	ret := C.get_string(&cstr)
	if ret != 0 {
		println("C function failed")
		os.Exit(1)
	}

	goStr := C.GoString(cstr)
	fmt.Println(goStr)

	C.free(unsafe.Pointer(cstr))
}
