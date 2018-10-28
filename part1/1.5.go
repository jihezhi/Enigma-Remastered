package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, e := os.Open("../README.md")
	if e != nil {
		fmt.Println(e.Error())
	}
	ss := strings.NewReader("reader?")
	bb := bytes.NewReader([]byte{'F', 'F', 'F'})
	k := []io.Reader{io.Reader(ss), bb, f}
	for _, v := range k {
		io.Copy(os.Stdout, v)
	}
}
