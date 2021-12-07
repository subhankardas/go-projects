package main

import "testing"

func Repeat(chr string, repeat int) string {
	var repeated string
	for count := 1; count <= repeat; count++ {
		repeated += chr
	}
	return repeated
}

/* Test a loop */
func TestRepeat(test *testing.T) {
	expected := "aaaa"
	actual := Repeat("a", 4)
	if actual != expected {
		test.Errorf("expected %q but got %q", expected, actual)
	}
}

/*
 Benchmarking - measures how long it takes to run a code N no. of times.
 Command: go test -bench=.
 Result : 10000000           136 ns/op
 136 ns/op means is our function takes on average 136 nanoseconds to run on local.

 Coverage: go test -cover
*/
func BenchmarkRepeat(bench *testing.B) {
	for count := 0; count < bench.N; count++ {
		Repeat("a", 4)
	}
}
