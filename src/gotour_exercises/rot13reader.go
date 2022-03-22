package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)
	for x := range b {
		char := b[x]
		if char >= 65 && char <= 90 {
			char += byte(13)
			if char > 90 {
				char -= 26
			}
		} else if char >= 97 && char <= 122 {
			char += 13
			if char > 122 {
				char -= 26
			}
		}
		b[x] = char
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
