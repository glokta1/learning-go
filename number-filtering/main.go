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

func main() {
}
