package array

// Sum some
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number

	}
	return sum
}

// SumAll sum of arrays
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails sum of arrays
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tails := numbers[1:]
		sums = append(sums, Sum(tails))
	}
	return sums
}
