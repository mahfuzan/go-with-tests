package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("add numbers", func(t *testing.T) {
		got := Add(5, 2)
		want := 7

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSubtract(t *testing.T) {
	t.Run("subtract numbers positive", func(t *testing.T) {
		got := Subtract(5, 3)
		want := 2

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("subtract numbers negative", func(t *testing.T) {
		got := Subtract(3, 5)
		want := -2

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("multiply numbers", func(t *testing.T) {
		got := Multiply(2, 3)
		want := 6

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestDivide(t *testing.T) {
	t.Run("divide numbers", func(t *testing.T) {
		got := Divide(2, 4)
		want := 0.5

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}
