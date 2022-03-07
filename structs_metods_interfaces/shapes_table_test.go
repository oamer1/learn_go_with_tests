// Identical to shapes_test but using table

// bug is found with Area it is very easy to add a new test case to exercise it before fixing it.
package main

import "testing"

func TestPerimeter1(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	// float64 and the .2 means print 2 decimal places
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea1(t *testing.T) {

	// "anonymous struct"
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
