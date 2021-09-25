package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 13, 6, 8}
	target := 12
	fmt.Println(getResult(nums, target))
}

func getResult(nums []int, target int) []int {
	answer := [2]int{0, 0}

	for i, item := range nums {
		for j, item2 := range nums {
			if nums[i] != nums[j] {
				if item+item2 == target {
					fmt.Println(item)
					return answer[i:j]
				}
			}
		}
	}

	return answer[:]
}
