package basics

import (
	"fmt"
	"sort"
)

/* Sorting */
func Basics10() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:   ", ints)

	sorted := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ", sorted)
}

/* Sorting by Functions - sort a collection by something other than its natural order. */
type fruits []string

// We implement sort.Interface - Len, Less, and Swap - on our type
// so we can use the sort package’s generic Sort function.
// Len and Swap will usually be similar across types.
func (fr fruits) Len() int {
	return len(fr)
}

func (fr fruits) Swap(i, j int) {
	fr[i], fr[j] = fr[j], fr[i]
}

// Less will hold the actual custom sorting logic.
func (fr fruits) Less(i, j int) bool {
	return len(fr[i]) < len(fr[j])
}

func Basics11() {
	fr := []string{"peach", "banana", "kiwi"}

	sort.Sort(fruits(fr)) // Use sort.Sort on our typed slice
	fmt.Println("fruits sorted by length", fr)
}
