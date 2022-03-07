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
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}

// The %#v format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.
