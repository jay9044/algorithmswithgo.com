package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + Sum(numbers[1:])
}

// THIS WORKS TOO
// func Sum(numbers []int) int {
// 	runningTotal := 0
// 	for _, num := range numbers {
// 		runningTotal += num
// 	}
// 	return runningTotal
// }
