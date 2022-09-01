package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat string", func(t *testing.T) {
		got := Repeat("shark", 3)
		want := "sharksharkshark"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("shark", 3)
	}
}
