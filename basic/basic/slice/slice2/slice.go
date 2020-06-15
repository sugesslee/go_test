package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)
	// len=0 cap=0 slice=[]

	numbers = append(numbers, 0)
	printSlice(numbers)
	// len=1 cap=1 slice=[0]

	numbers = append(numbers, 1)
	printSlice(numbers)
	// len=2 cap=2 slice=[0 1]

	numbers = append(numbers, 2, 3, 4, 5)
	printSlice(numbers)
	// len=6 cap=6 slice=[0 1 2 3 4 5]

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
	// len=6 cap=12 slice=[0 1 2 3 4 5]

	str1 := []string{"a", "b"}
	str2 := []string{"c"}
	str1 = append(str1, str2...)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(str1), cap(str1), str1)
	// abc
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
