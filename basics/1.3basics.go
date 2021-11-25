package main

import (
	"fmt"
	"strings"
)

func main() {
	/* Strings */
	var name string = "hello go!"
	len := len(name)                                    // Length of string
	var str = strings.Join([]string{"abc", "xyz"}, "-") // Join two strings

	fmt.Println("length of string", name, "is", len)
	fmt.Printf("second character of %v is %c\n", name, name[1])
	fmt.Println("concat abc and xyz is", str)

	/* Switch */
	switch name {
	case "hello go!":
		fmt.Println("got matching case!")
	default:
		fmt.Println("default case!")
	}

	whatamI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i'm a bool")
		case int:
			fmt.Println("i'm an int")
		default:
			fmt.Printf("don't know type %T\n", t)
		}
	}

	whatamI(true)
	whatamI(1)
	whatamI("hey")

	/* Arrays */
	var arr [5]int // Not initialized
	fmt.Println(arr)

	arr[2] = 4 // Set value at index
	fmt.Println(arr)

	arrn := [5]int{1, 2, 3} // Partially initialized
	fmt.Println(arrn)

	arrt := []int{1, 2, 3} // Fully initialized, inferred length
	fmt.Println(arrt)

	var arr2d [3][3]int // 2D array
	fmt.Println(arr2d)

	/* Slices - dynamic arrays */
	var nums1 []int // Undefined length, nil slice - no elements
	fmt.Println("slice", nums1, "is nil", nums1 == nil)
	printSlice(nums1)

	nums2 := []int{0, 1, 2, 3, 4, 5} // Slice initialization
	fmt.Println("slice", nums2, "is nil", nums2 == nil)
	printSlice(nums2)

	nums3 := make([]int, 3, 5) // Make slice of length and capacity
	printSlice(nums3)

	nums4 := nums2[2:4] // Sub-slicing
	printSlice(nums4)

	copy(nums3, nums4) // Copy slice
	printSlice(nums3)

	nums3 = append(nums3, 1, 2)
	printSlice(nums3)
}

func printSlice(nums []int) {
	fmt.Println("slice", nums, "length is", len(nums), "and capacity is", cap(nums))
}
