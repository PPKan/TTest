package main

func SumAllTails(s1, s2 []int) (SumAll []int) {

	// var SumAll []int

	sums := 0
	for _, number := range s1[1:] {
		sums = sums + number
	}
	SumAll = append(SumAll, sums)
	sums = 0
	for _, number := range s2[1:] {
		sums = sums + number
	}
	SumAll = append(SumAll, sums)

	return SumAll
}
