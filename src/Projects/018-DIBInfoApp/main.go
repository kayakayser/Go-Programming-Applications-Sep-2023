package main

import (
	"018-DIBInfoApp/csd/err"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		err.ExitFailure("wrong number of arguments!...")
	}

	f, e := os.Open(os.Args[1])

	if e != nil {
		err.ExitFailureError("Open:", e)
	}

	defer func() { _ = f.Close() }()

	var width int32
	var height int32

	_, e = f.Seek(18, io.SeekStart)
	if e != nil {
		err.ExitFailureError("Seek", e)
	}

	e = binary.Read(f, binary.LittleEndian, &width)
	if e != nil {
		err.ExitFailureError("Read", e)
	}

	if e != nil {
		err.ExitFailureError("Seek", e)
	}

	e = binary.Read(f, binary.LittleEndian, &height)
	if e != nil {
		err.ExitFailureError("Read", e)
	}

	fmt.Printf("Width = %d, Height = %d\n", width, height)
}
