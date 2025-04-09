package main

import (
	"reflect"
	"testing"
)


func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}
		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("got %d, expected: %d, given: %v", sum, expected, numbers)
		}
	})
}
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("sum tails", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("sum tails handles empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0, 9}, []int{})
		want := []int{2, 9, 0}

		checkSums(t, got, want)
	})

}
