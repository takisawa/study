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
	fmt.Printf("----- reflect.TypeOf -----\n")
	fmt.Println(reflect.TypeOf(3))
	fmt.Println(reflect.TypeOf("ABC"))
	fmt.Println(reflect.TypeOf([]int{1, 2, 3}))
	fmt.Println(reflect.TypeOf(map[string]int{"A": 1, "B": 2, "C": 3}))
	fmt.Println(reflect.TypeOf(Person{}))
	fmt.Println(reflect.TypeOf(&Person{}))
	fmt.Println(reflect.TypeOf(nil))
	fmt.Printf("%T\n", &Person{})

	fmt.Printf("----- reflect.ValueOf -----\n")
	fmt.Println(reflect.ValueOf(3))
	fmt.Println(reflect.ValueOf("ABC"))
	fmt.Println(reflect.ValueOf([]int{1, 2, 3}))
	fmt.Println(reflect.ValueOf(map[string]int{"A": 1, "B": 2, "C": 3}))
	fmt.Println(reflect.ValueOf(Person{}))
	fmt.Println(reflect.ValueOf(&Person{}))
	fmt.Println(reflect.ValueOf(nil))

	// reflect.New
	person := &Person{}
	personType := reflect.TypeOf(person)
	reflectedType := reflect.New(personType)

	fmt.Printf("person: %#v\n", person)
	fmt.Printf("personType: %#v\n", personType)
	fmt.Printf("reflectedType: %#v\n", reflectedType)
}
