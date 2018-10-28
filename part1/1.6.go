package main

import (
	"fmt"
	"sort"
)

func main() {
	f := map[float64]string{
		3.14159: "pi",
		2.714:   "e",
		1.414:   "sqrt2"}
	keys := make([]float64, 0)
	values := make([]string, 0)

	for k := range f {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for _, k := range keys {
		values = append(values, f[k])
	}
	fmt.Println(keys)
	fmt.Println(values)
}
