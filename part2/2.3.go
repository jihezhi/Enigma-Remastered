package main

import "fmt"

type player interface {
	eval() int
	describe() string
}

type me struct{}

func (v me) eval() int {
	return 5
}
func (v me) describe() string {
	return "nothing special"
}

type misaka struct{}

func (v misaka) eval() int {
	return 10000
}
func (v misaka) describe() string {
	return "Bilibili~"
}

type kirito struct{}

func (v kirito) eval() int {
	return 1000
}
func (v kirito) describe() string {
	return "The Black Swordsman"
}

func list(args ...player) {
	for _, v := range args {
		fmt.Println(v.eval())
		fmt.Println(v.describe())
	}
}

func main() {
	list(me{}, misaka{}, kirito{})
}
