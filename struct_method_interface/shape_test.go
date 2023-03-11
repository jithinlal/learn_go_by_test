package structmethodinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 10.0, height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	//checkArea := func(tb testing.TB, shape Shape, want float64) {
	//	tb.Helper()
	//	got := shape.Area()
	//
	//	if got != want {
	//		t.Errorf("got %.2f want %.2f", got, want)
	//	}
	//}
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{width: 12.0, height: 6.0}
	//	checkArea(t, rectangle, 72.0)
	//})
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//})

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12.0, height: 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12.0, height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})

	}
}
