package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNUmbers := len(numbersToSum)
	// sums := make([]int, lengthOfNUmbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
