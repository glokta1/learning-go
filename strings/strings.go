package main

import (
	"fmt"
	"unicode/utf8"
)

func getDeets(s string) {
	fmt.Println(s)

	// print each byte
	fmt.Printf("% x\n", s)

	// print total number of bytes in string
	fmt.Println("Number of bytes in string: ", len(s))

	// print number of "code points" in string
	fmt.Println("Number of code points/runes: ", utf8.RuneCountInString(s))
	fmt.Println()
}

func main() {
	s1 := "Tschçss"
	s2 := "für"
	s3 := "बेटी"
	s4 := "बेटा"

	getDeets(s1)
	getDeets(s2)
	getDeets(s3)
	getDeets(s4)

	var ns string
	for _, ch := range s1 {
		ns += string(ch)
	}

	fmt.Println(ns)
}
