package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// TODO: implement Read() method for rot13Reader
func (r rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	for i := 0; i < n; i++ {
		// ['Aa' ~ 'Mm'] to ['Nn' ~ 'Zz']
		if ('A' <= b[i] && b[i] <= 'M') || ('a' <= b[i] && b[i] <= 'm') {
			b[i] += 13
		// ['Nn' ~ 'Zz'] to ['Aa' ~ 'Mm']
		} else if ('N' <= b[i] && b[i] <= 'Z') || ('n' <= b[i] && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
