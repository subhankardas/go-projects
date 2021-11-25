package main

import "fmt"

func main() {
	/* Maps */
	mp := make(map[string]int)
	fmt.Println(mp)

	nmap := map[string]int{"usr1": 123, "usr2": 456} // Initialize a map
	fmt.Println(nmap, nmap["usr2"])

	mp["user1"] = 123 // Set key-value pair

	fmt.Println(mp["user1"]) // Get value using key
	fmt.Println(mp["user2"]) // No key found, returns 0
	fmt.Println(mp)

	fmt.Println("length of map is", len(mp))

	// Check key presence in map
	_, present := mp["user1"]
	fmt.Println("key user1 is present in map", present)

	delete(nmap, "usr1") // Delete a key-value pair
	fmt.Println(nmap)

	/* Range */
	nums := [3]int{1, 2, 3}
	for idx := range nums {
		fmt.Print(nums[idx], " ")
	}
	fmt.Println()

	nmap["usr1"] = 123
	for key, val := range nmap { // Read key-values in map
		fmt.Println(key, "->", val)
	}

	for idx, char := range "hello" { // Read characters in string
		fmt.Printf("%v[%c] ", idx, char)
	}

	fmt.Println()
	fmt.Println("sum 1+2 is", add(1, 2), "sum 1+2+3 is", add(1, 2, 3)) // Using variadic function

	/* Recursion */
	fmt.Println("fibonacci of 5 is", fibonacci(5))
}

/* Variadic function using range */
func add(nums ...int) int {
	var sum = 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

/* Recursive function */
func fibonacci(n int) int {
	if n == 0 {
		return 1
	}
	return n * fibonacci(n-1)
}
