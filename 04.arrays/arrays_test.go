package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	firstSum := []int{1, 2, 3}
	secondSum := []int{4, 5, 6}

	got := SumAll(firstSum, secondSum)
	want := []int{6, 15}
	assertArraysEqual(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("sum all numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		assertArraysEqual(t, got, want)
	})

	t.Run("safely sum an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}
		assertArraysEqual(t, got, want)
	})
}

func TestSumWithoutHead(t *testing.T) {

	t.Run("sum without the head of the slices", func(t *testing.T) {
		got := SumWithoutHead([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{3, 9}
		assertArraysEqual(t, got, want)
	})

	t.Run("safely sums with an empty slice", func(t *testing.T) {
		got := SumWithoutHead([]int{}, []int{4, 5, 6})
		want := []int{0, 9}
		assertArraysEqual(t, got, want)
	})

}

func assertArraysEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
