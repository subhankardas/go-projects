package main

import "fmt"

/* Structs are collection of fields */
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"sam doe", 12})             // Directly initialize fields
	fmt.Println(person{name: "mary doe", age: 23}) // Initialize by field names
	fmt.Println(person{name: "some doe"})          // Partially initialize

	per := newPerson("john doe") // Construct a new person
	fmt.Println(per, "age is", per.age)

	perp := &per // Pointer to the struct
	fmt.Println(perp, "age is", perp.age)
}

func newPerson(name string) person {
	person := person{name: name, age: 0}
	person.age = 20 // Structs are mutable
	return person
}
