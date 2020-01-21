package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "shape", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(tt.shape, tt.want)
		})
	}

}
