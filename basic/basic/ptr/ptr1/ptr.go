package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000
	// 指针ptr地址
	ptr = &a
	// 指向指针 ptr 地址
	pptr = &ptr

	fmt.Printf("a = %d\n", a)
	fmt.Printf("*ptr = %d\n", *ptr)
	fmt.Printf("&ptr = %d\n", &ptr)
	fmt.Printf("**ptr = %d\n", **pptr)
	fmt.Printf("&ptr = %d\n", &pptr)
}
