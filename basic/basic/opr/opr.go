package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	//a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	var aa int = 4
	var bb int32
	var cc float32
	var ptr *int

	fmt.Printf("aa 的类型 = %T\n", aa)
	fmt.Printf("bb 的类型 = %T\n", bb)
	fmt.Printf("cc 的类型 = %T\n", cc)

	ptr = &aa
	fmt.Printf("aa 的值为 %d\n", aa)
	fmt.Printf("ptr 为 %d\n", *ptr)

	var dd = 10
	a = dd
	ptr = &a
	fmt.Printf("aa 的值为 %d\n", aa)
	fmt.Printf("ptr 为 %d\n", *ptr)
}
