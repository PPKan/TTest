package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [...]int{1, 2, 3, 4, 5}

	got := sum(numbers)
	want := 15

	if got != want {
		t.Errorf("from '%v' we have got '%d' but wanted '%d'", numbers, got, want)
	}

}
