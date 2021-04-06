package structs

import (
	"testing"
)

func TestPerimiter(t *testing.T) {
	rectange := Rectange{10.0, 10.0}
	got := rectange.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	cases := []struct {
		shape Shape
		want  float64
	}{{shape: Rectange{10.0, 10.0}, want: 100.0}, {shape: Circle{10.0}, want: 314.1592653589793}}

	for _, testCase := range cases {
		shape := testCase.shape
		got := shape.Area()
		want := testCase.want

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
}
