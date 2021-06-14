package main

import "testing"

func TestSumAllTails(t *testing.T) {

	t.Run("check if two same work", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		for i := 0; i < len(got); i++ {
			if got[0] != want[0] {
				t.Errorf("got '%v' but want '%v'", got, want)
			}
		}
	})

	t.Run("check if different length work", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{3, 4})
		want := []int{5, 4}

		for i := 0; i < len(got); i++ {
			if got[0] != want[0] {
				t.Errorf("got '%v' but want '%v'", got, want)
			}
		}
	})
}
