package main

import "fmt"

// 匿名函数的优越性在于可以直接使用函数内的变量，不必声明
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 闭包使用
func add(x1, x2 int) func(x3, x4 int) (int, int, int) {
	i := 0
	return func(x3, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}

func main() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	add_func := add(1, 2)
	fmt.Println(add_func(1, 1))
	fmt.Println(add_func(0, 0))
	fmt.Println(add_func(2, 2))
}
