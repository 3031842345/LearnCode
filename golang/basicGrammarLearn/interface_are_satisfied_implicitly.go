package main

import (
	"fmt"
	"os"
)

//Package io defines Reader、 Writer and ReadWriter;
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w ReadWriter

	w = os.Stdout

	fmt.Fprintf(w, "hello, interface!\n")
}
