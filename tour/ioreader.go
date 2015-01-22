package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(res []byte) (int, error) {
	data := make([]byte, 100)
	r.r.Read(data)
	for i, _ := range data {
		data[i] = data[i] + 1
		res[i] = data[i]
	}
	return len(data), nil
}
func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
