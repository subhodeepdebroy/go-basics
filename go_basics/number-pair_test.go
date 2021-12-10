package go_basics

import (
	"testing"
)

func TestFindPair(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got, _ := FindPair(input, 3)
	want := [2]int{1, 2}
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkFindPair(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	for n := 0; n < b.N; n++ {
		FindPair(input, 3)
	}
}
