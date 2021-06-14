package main

import (
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{4, 6}

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got '%v' but want '%v'", got, want)
	// }

	for i := 0; i < len(got); i++ {
		if got[0] != want[0] {
			t.Errorf("got '%v' but want '%v'", got, want)
		}
	}
}
