package arrays

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for _, toSum := range numbers {
		sums = append(sums, Sum(toSum))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, toSum := range numbers {
		if len(toSum) == 0 {
			sums = append(sums, 0)
		} else {
			tail := toSum[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func SumWithoutHead(numbers ...[]int) []int {
	var sums []int

	for _, toSum := range numbers {
		if len(toSum) == 0 {
			sums = append(sums, 0)
		} else {
			withoutHead := toSum[:len(toSum)-1]
			sums = append(sums, Sum(withoutHead))
		}
	}

	return sums
}
