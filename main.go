package main

import "fmt"

type MyElement struct {
	x int
}

func (e *MyElement) Get() int {
	return e.x
}

func main() {
	var a uint64 = 1
	fmt.Println(foo() == a)
}

func foo() interface{} {
	return 1
}
