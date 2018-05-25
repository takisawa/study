package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func main() {
	// reflect.Typeof
	fmt.Println(reflect.TypeOf(3))
	fmt.Println(reflect.TypeOf("ABC"))
	fmt.Println(reflect.TypeOf([]int{1, 2, 3}))
	fmt.Println(reflect.TypeOf(map[string]int{"A": 1, "B": 2, "C": 3}))
	fmt.Println(reflect.TypeOf(Person{}))
	fmt.Println(reflect.TypeOf(&Person{}))

	// reflect.New
	person := &Person{}
	personType := reflect.TypeOf(person)
	reflectedType := reflect.New(personType)

	fmt.Printf("person: %#v\n", person)
	fmt.Printf("personType: %#v\n", personType)
	fmt.Printf("reflectedType: %#v\n", reflectedType)
}
