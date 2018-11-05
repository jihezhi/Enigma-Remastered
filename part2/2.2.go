package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SquareMatrix struct {
	dim    int
	values [][]float32
}

func buildRandom(n int) SquareMatrix {
	v := make([][]float32, n)
	for row := range v {
		v[row] = make([]float32, n)
		for inner := range v[row] {
			v[row][inner] = rand.Float32()
		}
	}
	result := SquareMatrix{dim: n, values: v}
	return result
}

func buildRandomInt(n int) SquareMatrix {
	v := make([][]float32, n)
	for row := range v {
		v[row] = make([]float32, n)
		for inner := range v[row] {
			rand.Seed(time.Now().UnixNano())
			v[row][inner] = float32(rand.Intn(10))
		}
	}
	result := SquareMatrix{dim: n, values: v}
	return result
}

func (it SquareMatrix) remove1stRowAndColumn(colIdx int) SquareMatrix {
	if colIdx >= it.dim {
		panic("what have you done???")
	} else {
		v := make([][]float32, it.dim-1)
		for i := range v {
			v[i] = make([]float32, it.dim-1)
			for j := range v[i] {
				if j < colIdx {
					v[i][j] = it.values[i+1][j]
				} else {
					v[i][j] = it.values[i+1][j+1]
				}
			}
		}
		return SquareMatrix{dim: it.dim - 1, values: v}
	}
}

func (it SquareMatrix) det() float32 {
	var sum float32
	dim := it.dim
	// fmt.Println(dim, it.values)
	if dim == 0 {
		return 1
	} else if dim == 1 {
		return it.values[0][0]
	} else {
		for i := range it.values {
			sign := (i%2)*(-2) + 1
			sum += float32(sign) * it.values[0][i] * it.remove1stRowAndColumn(i).det()
			// fmt.Println(sign, sum)
		}
		return sum
	}
}

func simpleTest() {
	test := buildRandomInt(2)
	fmt.Println(test.values)
	fmt.Println(test.dim)
	fmt.Println(test.det())
}

func worker(dim int, out chan float32) {
	test := buildRandom(dim)
	result := test.det()
	out <- result
}

func workerHelper(mat SquareMatrix, out chan float32) {
	result := mat.det()
	out <- result
}

func main() {
	var sum1 float32
	var sum2 float32
	mats := make([]SquareMatrix, 10000)

	for i := 0; i < 10000; i++ {
		mats[i] = buildRandom(6)
	}

	start := time.Now()
	for i := 0; i < 10000; i++ {
		sum1 += mats[i].det()
	}
	elapse := time.Since(start)
	fmt.Println(sum1, elapse)

	tube := make(chan float32, 10000)
	q := make(chan int)
	start = time.Now()
	for i := 0; i < 10000; i++ {
		select {
		case q <- 0:
			go workerHelper(mats[i], tube)
			<-q
		default:
			workerHelper(mats[i], tube)
		}
	}
	for i := 0; i < 10000; i++ {
		sum2 += <-tube
	}
	elapse = time.Since(start)
	fmt.Println(sum2, elapse)
}
