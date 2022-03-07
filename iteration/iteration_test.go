package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

func BenchmarkRepeat(b *testing.B) {
	// t runs b.N times and measures how long it takes.
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	rep := Repeat("n")
	fmt.Println(rep)
	// Output:nnnnn
}
