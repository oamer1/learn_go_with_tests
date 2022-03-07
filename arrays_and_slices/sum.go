package main

// An interesting property of arrays is that the size is encoded in its type.
func Sum(numbers []int) int {
	sum := 0
	// _ is blank identifier
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Go can let you write variadic functions that can take a variable number of arguments.
func SumAll(numsToSum ...[]int) []int {
	lengthOfNumbers := len(numsToSum)
	sums := make([]int, lengthOfNumbers)

	for i, nums := range numsToSum {
		sums[i] = Sum(nums)
	}
	return sums
}
