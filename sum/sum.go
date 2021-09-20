package main

func Sum(n []int) int {

	sum := 0

	//for i := 0; i < len(n); i++ {
	//	sum += n[i]
	//}

	for _, number := range n {
		sum += number
	}

	return sum
}

func SumAll(n ...[]int) []int {
	var sums []int

	for _, numbers := range n {
		sums = append(sums, Sum(numbers))
	}

	//sums := make([]int, len(n))
	//for i, numbers := range n {
	//	sums[i] = Sum(numbers)
	//}

	return sums
}

func SumAllTails(n ...[]int) []int {
	var sums []int
	for _, numbers := range n {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
