package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Test 1,2,3,4,5", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15

		if got != want {
			t.Errorf("got %d != want %d", got, want)
		}
	})
	t.Run("Test 1,2,3", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		want := 6

		if got != want {
			t.Errorf("got %d != want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v != want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v != want %v", got, want)
	}
}
