package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 5
	marks := [3]int{23, 45, 67}

	/* Loop execution */
	for idx := 0; idx < 5; idx++ {
		fmt.Print(idx+1, " ")
	}

	for num >= 1 {
		fmt.Print(num, " ")
		num--
	}

	for idx := range marks {
		fmt.Print(marks[idx], " ")
	}

	/* Functions */
	var res int = max(34, 56)
	fmt.Println("\nmax is", res)

	var str1, str2 = swap("abc", "xyz") // Return multiple values
	fmt.Println("swapped abc xyz is", str1, str2)

	/* Function call */
	var n1, n2 = 11, 20
	fmt.Println("n1 is", n1, "n2 is", n2)

	addbyvalue(n1, n2) // Call by value - default
	fmt.Println("n1 is", n1, "n2 is", n2)

	addbyref(&n1, &n2) // Call by reference
	fmt.Println("n1 is", n1, "n2 is", n2)

	// Function as a value
	square := func(num int) int { return num * num }
	fmt.Println("square of 23 is", square(23))

	// Function closure
	next := incrementer()
	fmt.Println(next(), next(), next())

	// Using methods
	circle := Circle{x: 0, y: 0, radius: 12}
	fmt.Println("area of circle is", circle.area())
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

/* Public function names starts with uppercase */
func PublicFunc() {
	fmt.Println("this is a public function")
}

/* Private function names starts with lowercase */
func privateFunc() {
	fmt.Println("this is a private function")
}

/* Function returning multiple values */
func swap(name1 string, name2 string) (string, string) {
	return name2, name1
}

/* Function with call by value */
func addbyvalue(num1 int, num2 int) int {
	num1 += 1
	num2 += 1
	return num1 + num2
}

/* Function with call by reference */
func addbyref(num1 *int, num2 *int) int {
	*num1 += 1
	*num2 += 1
	return *num1 + *num2
}

/* Function closures */
func incrementer() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/* Methods */
type Circle struct {
	// Define a circle
	x, y, radius float64
}

// Define a method for circle
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
