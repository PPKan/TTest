package main

func SumAll(s1 []int, s2 []int) []int {
	// SumList = make([]int, len(s1))
	var SumList []int
	// if len(s1) != len(s2) {
	// 	return []int{0,0}
	// }
	for i := 0; i < len(s1); i++ {
		// SumList[i] = s1[i] + s2[i]
		number := s1[i] + s2[i]
		SumList = append(SumList, number)
	}

	return SumList
}
