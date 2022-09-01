package main

import (
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("sum array of 4 numbers", func(t *testing.T) {
		got := SumArray([4]int{2, 2, 3, 5})
		want := 12
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
func TestSum(t *testing.T) {
	t.Run("sum slice of numbers", func(t *testing.T) {
		got := Sum([]int{2, 3, 5})
		want := 10
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum some of the slice", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{1, 2, 3}, []int{3, 5, 1, 8})
		want := []int{3, 5, 14}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("sum with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 2, 1})
		want := []int{0, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
