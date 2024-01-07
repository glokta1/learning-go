package main

import (
	"slices"
	"testing"
)

func TestExtractEven(t *testing.T) {
	got := ExtractEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{2, 4, 6, 8, 10}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
