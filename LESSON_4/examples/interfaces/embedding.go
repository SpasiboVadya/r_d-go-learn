package interfaces

import "fmt"

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	copy(p, "data-from-reader")
	return len(p), nil
}

type MyRW struct {
	MyReader // Embedded
}

func (rw MyRW) Write(p []byte) (int, error) {
	fmt.Println("Writing:", string(p))
	return len(p), nil
}

func RunRW() {
	var r Reader = MyReader{}
	buf := make([]byte, 16)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Println("Reader error:", err)
		return
	}
	fmt.Println("Reader output:", string(buf[:n]))

	var rw ReadWriter = MyRW{}
	n, err = rw.Read(buf)
	if err != nil {
		fmt.Println("Reader error:", err)
		return
	}
	fmt.Println("ReadWriter read:", string(buf[:n]))

	write, err := rw.Write([]byte("hello"))
	if err != nil {
		fmt.Println("Writer error:", err)
		return
	}
	fmt.Println(write)
}
