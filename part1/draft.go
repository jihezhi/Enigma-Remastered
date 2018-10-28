package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
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

	randArray := []int{11, 2, 3}
	sort.Ints(randArray)
	fmt.Println(randArray)

	fmt.Println(strconv.FormatInt(int64(age), 16))
	sqrt := "1.414"
	floatSqrt, _ := strconv.ParseFloat(sqrt, 64)
	fmt.Println(floatSqrt)

	f, _ := os.Open("README.md")
	io.Copy(os.Stdout, f)

	ss := strings.NewReader("reader?")
	bb := bytes.NewReader([]byte{'F', 'F', 'F'})
	k := []io.Reader{io.Reader(ss), bb, f}
	for _, v := range k {
		io.Copy(os.Stdout, v)
	}
}
