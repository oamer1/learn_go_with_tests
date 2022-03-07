// Arrays allow you to store multiple elements of the same type in a variable in a particular order.

package main

import "testing"

func TestSum(t *testing.T) {

	// Alternatively [...]int{1, 2, 3, 4, 5}
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	// %v placeholder to print the "default" format, which works well for arrays.
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
