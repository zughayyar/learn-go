package arrays

import "fmt"

func Sum(arr []int) int {
	var sum int
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, arr := range numbersToSum {
		sum = append(sum, Sum(arr))
	}
	return sum
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
