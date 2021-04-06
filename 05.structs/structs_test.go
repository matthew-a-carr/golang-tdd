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
	}{{Rectange{10.0, 10.0}, 100.0}, {Circle{10.0}, 314.1592653589793}}

	for _, testCase := range cases {
		shape := testCase.shape
		got := shape.Area()
		want := testCase.want

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectange := Rectange{10.0, 10.0}
		want := 100.0
		checkArea(t, rectange, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}
