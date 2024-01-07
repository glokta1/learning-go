package main

type check func(int) bool

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func isMultipleOfThree(n int) bool {
	return n%3 == 0
}

func isMultipleOfFive(n int) bool {
	return n%5 == 0
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n <= 1 || n > 3 && (n%2 == 0 || n%3 == 0) {
		return false
	}

	for i := 5; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func ExtractEven(s []int) []int {
	evenSlice := make([]int, 0, len(s))
	for _, num := range s {
		if isEven(num) {
			evenSlice = append(evenSlice, num)
		}
	}

	return evenSlice
}

func ExtractOdd(s []int) []int {
	oddSlice := make([]int, 0, len(s))
	for _, num := range s {
		if isOdd(num) {
			oddSlice = append(oddSlice, num)
		}
	}

	return oddSlice
}

func ExtractPrimes(s []int) []int {
	primesSlice := make([]int, 0, len(s))
	for _, num := range s {
		if isPrime(num) {
			primesSlice = append(primesSlice, num)
		}
	}

	return primesSlice
}

func ExtractOddPrimes(s []int) []int {
	return ExtractOdd(ExtractPrimes(s))
}

func ExtractEvenAndFives(s []int) []int {
	evenFives := make([]int, 0, len(s))
	for _, num := range s {
		if isMultipleOfFive(num) && isEven(num) {
			evenFives = append(evenFives, num)
		}
	}

	return evenFives
}

func ExtractOddThreesAboveTen(s []int) []int {
	oddThreesAboveTen := make([]int, 0, len(s))
	for _, num := range s {
		if num > 10 && isMultipleOfThree(num) && isOdd(num) {
			oddThreesAboveTen = append(oddThreesAboveTen, num)
		}
	}

	return oddThreesAboveTen
}

func ExtractMatchingInts(s []int, conditions ...check) []int {
	matches := make([]int, 0, len(s))
	for _, num := range s {
		all := true
		for _, condition := range conditions {
			if !condition(num) {
				all = false
				break
			}
		}

		if all {
			matches = append(matches, num)
		}
	}

	return matches
}

func ExtractMatchingIntsAny(s []int, conditions ...check) []int {
	matches := make([]int, 0, len(s))
	for _, num := range s {
		for _, condition := range conditions {
			if condition(num) {
				matches = append(matches, num)
				break
			}
		}
	}

	return matches
}
func main() {
}
