package main

import (
	"fmt"
	"runtime"
	"time"
)

type Any struct {
	Obj *Any
}

func finalizer(o *Any) {
	fmt.Printf("run finalizer for %#v\n", *o)
}

func main() {
	o1 := Any{}
	o2 := Any{}

	o1.Obj = &o2
	o2.Obj = &o1

	runtime.SetFinalizer(&o1, finalizer)
	runtime.SetFinalizer(&o2, finalizer)

	runtime.GC()

	time.Sleep(1 * time.Second)
}
