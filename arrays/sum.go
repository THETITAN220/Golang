package main

func Sum(number []int) int {
	sum := 0
	for _, digit := range number {
		sum += digit
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNum := len(numbersToSum)
// 	sum := make([]int, lengthOfNum)
//
// 	for i, numbers := range numbersToSum {
// 		sum[i] = Sum(numbers)
// 	}
// 	return sum
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
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
