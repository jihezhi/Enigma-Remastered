package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	s := "hello~> <"
	age := 17
	pi := 3.1415926
	check := [3]bool{true, false, true}
	e := errors.New("boom")

	io.WriteString(os.Stdout, s)
	io.WriteString(os.Stdout, string(age))
	io.WriteString(os.Stdout, strconv.FormatFloat(pi, 'f', 10, 32))
	io.WriteString(os.Stdout, fmt.Sprint(check))
	io.WriteString(os.Stdout, e.Error())
}
