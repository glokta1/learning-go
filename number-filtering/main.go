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

func main() {

}
