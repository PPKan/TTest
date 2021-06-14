package main

func SumAllTails(s1, s2 []int) (SumAll []int) {

	// var SumAll []int

	sums1 := 0
	sums2 := 0

	for _, number := range s1[1:] {
		sums1 = sums1 + number
	}
	for _, number := range s2[1:] {
		sums2 = sums2 + number
	}

	SumAll = append(SumAll, sums1, sums2)

	return SumAll
}
