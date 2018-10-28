package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 17

	fmt.Println(strconv.FormatInt(int64(age), 16))
	sqrt := "1.414"
	floatSqrt, _ := strconv.ParseFloat(sqrt, 64)
	fmt.Println(floatSqrt)
}
