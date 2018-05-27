package main

import (
	"fmt"
	"io"
	"os"
)

type WriterWrapper struct {
	writer io.Writer
	count  int64
}

func (w *WriterWrapper) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	w.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := WriterWrapper{writer: w}

	return &writer, &writer.count
}

func main() {
	writer, n := CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "Hello World!\n")
	fmt.Println(*n)

	fmt.Fprintf(writer, "Hello World!\n")
	fmt.Println(*n)
}
