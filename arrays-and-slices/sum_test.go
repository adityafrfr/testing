package arrayAndSlices

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want	{
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestAll(t *testing.T) {
	t.Run("sum of 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0,9})
		want := []int{3, 9}

		if !slices.Equal(got, want)	{
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("test SumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{2,3}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want)	{
			t.Errorf("got %v want %v", got, want)
		}
	})
}