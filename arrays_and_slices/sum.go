package main

// An interesting property of arrays is that the size is encoded in its type.
func Sum(numbers [5]int) int {
	sum := 0
	// _ is blank identifier
	for _, num := range numbers {
		sum += num
	}
	return sum
}
