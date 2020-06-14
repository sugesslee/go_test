package main

import "fmt"

var str = "1"

func main() {
	fmt.Println("Hello World！")

	var b bool = true
	fmt.Println(b)

	var a string = "q"

	fmt.Println(a)

	var x, y = 1, 2

	fmt.Println(x, y)

	var z int = 0
	fmt.Println(z)

	// nil类型
	var a1 * int
	var a2 [] int
	var a3 map[string] int
	var a4 chan int
	var a5 func(string) int
	var a6 error //error是接口
	fmt.Println(a1, a2, a3, a4, a5, a6)

	fmt.Println(&a)
}
