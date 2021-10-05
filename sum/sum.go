package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sum := make([]int, lengthOfNumbers)

	for i, number := range numbersToSum {
		sum[i] = Sum(number)
	}
	return sum
}

func SumAlltails(numbersToTail ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToTail {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
