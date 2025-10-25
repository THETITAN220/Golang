package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("Want %.2f got %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
//
// 	t.Run("testing the area of rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{4.0, 2.0}
// 		want := 8.0
// 		checkArea(t, rectangle, want)
// 	})
//
// 	t.Run("testing the area of circles", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 7}, hasArea: 41.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("for %#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
