package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	isUpper := x >= 'A' && x <= 'Z'
	isLower := x >= 'a' && x <= 'z'
	isLetter := isUpper || isLower
	if isLetter {
		base := 'A'
		if isLower {
			base = 'a'
		}

		return (x-byte(base)+13)%26 + byte(base)
	}

	// Not a letter
	return x
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	for i := range b {
		b[i] = rot13(b[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
