package main

import (
	"fmt"
	"reflect"
)

func main() {
	ary1 := []int{1, 2, 3}
	ary2 := []int{1, 2, 3}
	ary3 := []int{1, 2, 3, 4}
	ary4 := []int{1, 3, 2}

	fmt.Println(reflect.DeepEqual(ary1, ary2))
	fmt.Println(reflect.DeepEqual(ary1, ary3))
	fmt.Println(reflect.DeepEqual(ary1, ary4))
}
