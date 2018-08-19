package main

import (
	"fmt"
	"runtime"
	"time"
)

func finalizer(s *string) {
	fmt.Printf("run finalizer for %#v\n", *s)
}

func main() {
	s := "Hello World"
	runtime.SetFinalizer(&s, finalizer)

	runtime.GC()
	time.Sleep(1 * time.Second)
}
