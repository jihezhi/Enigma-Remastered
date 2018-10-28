package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, e := os.Open("../README.md")
	if e != nil {
		fmt.Println(e.Error())
	}
	io.Copy(os.Stdout, f)
}
