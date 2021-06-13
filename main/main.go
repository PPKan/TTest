package main

import "fmt"

func main() {
	twoSum([], 9)
}

func twoSum(nums []int, target int) []int {
	var s []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (nums[i]+nums[j] == target) && (i != j) {
				s = append(s, i, j)
				return s
			}
		}
	}
}
