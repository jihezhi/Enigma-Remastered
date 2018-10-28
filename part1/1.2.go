package main

import (
	"fmt"
	"sort"
)

func main() {

	randArray := []int{11, 2, 3}
	sort.Ints(randArray)
	fmt.Println(randArray)
}
