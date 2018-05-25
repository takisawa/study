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
	person := &Person{}
	personType := reflect.TypeOf(person)

	fmt.Printf("person: %#v\n", person)
	fmt.Printf("personType: %#v\n", personType)
	fmt.Printf("reflect.New(personType): %#v\n", reflect.New(personType))
}
