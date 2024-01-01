package main

import (
	"fmt"
	"math"
)

var (
	d bool
	e string = "This is a string"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func sqrt(x float64) float64 {
	z := 1.0
	for {
		z_new := z - ((z*z - x) / (2 * z))
		if math.Abs(z-z_new) < 1e5 {
			break
		}

		z -= z_new
		fmt.Println(z)
	}

	return z
}

func my_map(f func(int) int, v []int) []int {
	var ans []int
	for _, elem := range v {
		ans = append(ans, f(elem))
	}

	return ans
}

func square(x int) int {
	return x * x
}

type Point3D struct {
	x int
	y int
	z int
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		temp := a
		a, b = b, a+b
		return temp
	}
}

func distance(p Point3D) float64 {
	return math.Sqrt(float64(p.x)*float64(p.x) + float64(p.y)*float64(p.y) + float64(p.z)*float64(p.z))
}

func main() {
	// a, b, c := 3, 4, 5
	// fmt.Println(add(3, 5))
	// fmt.Println(swap("Rafey", "Ahmad"))
	// fmt.Println(a, b, c)
	// fmt.Println(d, e, f)
	// fmt.Println(sqrt(25))
	// fmt.Println(sqrt(16))
	// a := 3
	// p := &a
	// fmt.Println(p)
	// point := Point3D{3, 5, 9}
	// loc := &point
	// fmt.Println(loc)
	// loc.x = 1008
	// fmt.Println(point.x)
	// board := [][]string{{".", ".", "."}, {"X", ".", "X"}}
	// fmt.Println(board)
	// for i, row := range board {
	// 	fmt.Printf("Row %d: %s\n", i, strings.Join(row, " "))
	// }

	// m := map[string]Point3D{
	// 	"A": {1, 2, 3},
	// 	"B": {4, 5, 6},
	// 	"C": Point3D{7, 8, 9},
	// }
	// fmt.Println(m)
	// v := []int{1, 2, 3, 4, 5}
	// squared := my_map(sqrt, v)
	// fmt.Println(squared)
	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	// s1 := "abcd"
	s1 := "a😭"
	s2 := []byte("a😭")

	fmt.Println(s1[0], s1[1], s1[2], s1[3])
	fmt.Println(s2)

	fmt.Printf("% x\n", s1)

}
