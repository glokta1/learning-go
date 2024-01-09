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

func TestExtractOdd(t *testing.T) {
	got := ExtractOdd([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{1, 3, 5, 7, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestExtractPrimes(t *testing.T) {
	got := ExtractPrimes([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{2, 3, 5, 7}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestExtractOddPrimes(t *testing.T) {
	got := ExtractOddPrimes([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{3, 5, 7}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestExtractEvenAndFives(t *testing.T) {
	got := ExtractEvenAndFives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	want := []int{10, 20}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestExtractOddThreesAboveTen(t *testing.T) {
	got := ExtractOddThreesAboveTen([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	want := []int{15}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestExtractMatchingInts(t *testing.T) {
	got := ExtractMatchingInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, isEven)
	want := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestExtractMatchingInts_MultipleConditions(t *testing.T) {
	got := ExtractMatchingInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, isEven, isMultipleOfThree)
	want := []int{6, 12, 18}

	if !slices.Equal(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestExtractMatchingIntsAny(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		conditions []Condition
		want       []int
	}{
		{
			name:       "Empty input",
			input:      []int{},
			conditions: []Condition{isEven},
			want:       []int{},
		},
		{
			name:       "No matches",
			input:      []int{1, 3, 5, 7, 9},
			conditions: []Condition{isEven},
			want:       []int{},
		},
		{
			name:       "Single condition, multiple matches",
			input:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			conditions: []Condition{isEven},
			want:       []int{2, 4, 6, 8, 10},
		},
		{
			name:       "Multiple conditions, multiple matches",
			input:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			conditions: []Condition{isEven, isMultipleOfThree},
			want:       []int{2, 3, 4, 6, 8, 9, 10},
		},
		{
			name:       "Multiple conditions, no matches",
			input:      []int{1, 3, 7, 9},
			conditions: []Condition{isEven, isMultipleOfFive},
			want:       []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractMatchingIntsAny(tt.input, tt.conditions...)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
func TestCreateMultipleFunc(t *testing.T) {
	tests := []struct {
		name string
		n    int
		x    int
		want bool
	}{
		{
			name: "Multiple of 2",
			n:    2,
			x:    4,
			want: true,
		},
		{
			name: "Not multiple of 2",
			n:    2,
			x:    5,
			want: false,
		},
		{
			name: "Multiple of 3",
			n:    3,
			x:    9,
			want: true,
		},
		{
			name: "Not multiple of 3",
			n:    3,
			x:    10,
			want: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := CreateMultipleFunc(tt.n)
			got := f(tt.x)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
