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

	var sums []int
	// ppend function which takes a slice and a new value, then returns a new slice with all the items in it.
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
