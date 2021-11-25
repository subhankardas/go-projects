/*
Go is a cross-platform, open source programming language.

Go can be used to create high-performance applications.

Go is a fast, statically typed, compiled language that feels
like a dynamically typed, interpreted language.

Go's syntax is similar to C++.

It provides garbage collection, type safety, dynamic-typing
capability, many advanced built-in types such as variable length
arrays and key-value maps. It also provides a rich standard library.
*/

// 1. Package declaration
package main

// 2. Import packages
import "fmt"

var sum int = 10 // Global variable declaration

// 3. Functions
func main() {
	// 4. Variables
	var name string = "go!" // ; is added automatically in compile-time

	// 5. Statements and expressions
	fmt.Println("hello " + name)
	fmt.Printf("sum is %v (global variable)\n", sum)

	/* Static type declaration */
	var num int16 = 12
	/* Dynamic type declaration */
	sum := 4.4 // Local variable has more precedence over global

	fmt.Printf("num %v is of type %T and ", num, num)
	fmt.Printf("sum %v is of type %T \n", sum, sum)

	/* Mixed variable declaration */
	var age, id = 31, "user123"
	fmt.Printf("%v type %T and age is %v type %T \n", id, id, age, age)

	var (
		n1 int = 1
		n2 int = 2
	)
	fmt.Println("sum is", n1+n2)

	/* Constants / literals */
	const PI = 3.14
	const (
		HEIGHT int = 5
		SHAPE      = "square"
	)

	fmt.Println("value of PI is", PI)
	fmt.Println("shape is", SHAPE, "height is", HEIGHT)
}
