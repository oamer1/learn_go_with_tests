package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Examples are compiled (and optionally executed) as part of a package's test suite.
// By adding this code the example will appear in the documentation inside godoc, making your code even more accessible.
func ExampleAdd() {
	sum := Add(3, 3)
	fmt.Println(sum)
	// Output:6
}
