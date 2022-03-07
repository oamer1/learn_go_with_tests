package main

func Sum(numbers [5]int) int {
	sum := 0
	// _ is blank identifier
	for _, num := range numbers {
		sum += num
	}
	return sum
}
