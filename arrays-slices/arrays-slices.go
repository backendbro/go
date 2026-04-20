package main

func Sum1(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}

// 	return sum
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// func Sum2(numbers []int) int {
// 	sum := 0
// 	for _, number := range numbers {
// 		sum += number
// 	}

// 	return sum
// }

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

// func SumAll1(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sum := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {

// 		sum[i] = Sum(numbers)
// 	}

// 	return sum
// }

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}

	return sum
}

// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sum []int

// 	for _, numbers := range numbersToSum {
// 		tails := numbers[1:]
// 		sum = append(sum, Sum(tails))
// 	}

// 	return sum
// }

func SumAllAppend(numbersToSum ...[]int) []int {
	var sum []int

	for _, number := range numbersToSum {
		sum = append(sum, Sum(number))
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tails := numbers[1:]
			sum = append(sum, Sum(tails))
		}
	}

	return sum
}
