package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{height: 12, width: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{height: 12, base: 6}, hasArea: 36.1},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

//func TestArea(t *testing.T) {
//
//	checkArea := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g, want %g", got, want)
//		}
//	}
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{10.0, 10.0}
//		checkArea(t, rectangle, 100.0)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10.0}
//		checkArea(t, circle, 314.1592653589793)
//	})
//}
