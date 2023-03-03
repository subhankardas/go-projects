package advanced

import (
	"fmt"
	"math"
)

/*
Generics - Generics allows us to create a single class/interface/method that can be
used with different types of data. The main aim of generics is to achieve greater flexibility
in terms of writing code with the addition of fewer lines.
*/
func Advanced15() {
	var r1 uint32 = 2
	var r2 float32 = 4.3

	fmt.Println("area of r1:", area(r1))
	fmt.Println("area of r2:", area(r2))

	fmt.Println("  2 + 3   :", add(2, 3))
	fmt.Println("3.1 + 2.4 :", add(3.1, 2.4))
	fmt.Println("ab  + cd  :", add("ab", "cd"))
}

// Approach 1: Using typed parameters to define custom data type.
func area[T uint32 | float32](r T) T {
	switch r := interface{}(r).(type) { // switch operations based on data type
	case uint32:
		return T(3 * r * r)
	case float32:
		return T(math.Pi * r * r)
	default:
		panic("unsupported type")
	}
}

// Approach 2: Using parameterized interface types
type T interface {
	int | float64 | string
}

func add[mytype T](val1, val2 mytype) mytype {
	return val1 + val2
}
