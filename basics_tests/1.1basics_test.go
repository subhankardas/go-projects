package main

import (
	"testing"
)

func Hello(name string) string {
	return "hello " + name + "!"
}

/* Simple unit test function */
func TestHello(test *testing.T) {
	want := "hello world!" // Given
	got := Hello("world")  // When
	if got != want {       // Then
		test.Errorf("got %q want %q", got, want)
	}
}

/* Efficient unit test function with multiple test cases */
func TestHelloAll(test *testing.T) {
	assert := func(test *testing.T, got, want string) {
		test.Helper()
		if got != want {
			test.Errorf("got %q want %q", got, want)
		}
	}

	test.Run("say hello to john", func(test *testing.T) {
		want := "hello john!"   // Given
		got := Hello("john")    // When
		assert(test, got, want) // Then
	})

	test.Run("say hello to no one", func(test *testing.T) {
		want := "hello !"       // Given
		got := Hello("")        // When
		assert(test, got, want) // Then
	})
}

func Add(x int, y int) int {
	return x + y
}

func TestAdd(test *testing.T) {
	test.Run("test add 2+3", func(test *testing.T) {
		expected := 5
		got := Add(2, 3)
		if got != expected {
			test.Errorf("expected %q but got %q", expected, got)
		}
	})
	test.Run("test add 6+9", func(test *testing.T) {
		expected := 15
		got := Add(6, 9)
		if got != expected {
			test.Errorf("expected %q but got %q", expected, got)
		}
	})
}
