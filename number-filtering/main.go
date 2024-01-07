package main

func ExtractEven(s []int) []int {
	evenSlice := make([]int, 0, len(s))
	for _, num := range s {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}

	return evenSlice
}

func ExtractOdd(s []int) []int {
	oddSlice := make([]int, 0, len(s))
	for _, num := range s {
		if num%2 != 0 {
			oddSlice = append(oddSlice, num)
		}
	}

	return oddSlice
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
	fives := make([]int, 0, len(s))
	for _, num := range s {
		if num%5 == 0 {
			fives = append(fives, num)
		}
	}

	return ExtractEven(fives)
}

func ExtractOddThreesAboveTen(s []int) []int {
	threesAboveTen := make([]int, 0, len(s))
	for _, num := range s {
		if num > 10 && num%3 == 0 {
			threesAboveTen = append(threesAboveTen, num)
		}
	}

	return ExtractOdd(threesAboveTen)
}

func main() {
}
