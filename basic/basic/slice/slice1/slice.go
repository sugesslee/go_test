package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)

	fmt.Println("number[1:4] = ", numbers[1:4])
	printSlice(numbers)
	// len=10 cap=10 slice=[0 1 2 3 4 5 6 7 8 9]
	// number[1:4] =  [1 2 3]
	// len=10 cap=10 slice=[0 1 2 3 4 5 6 7 8 9]

	fmt.Println("number[:3] = ", numbers[:3])
	printSlice(numbers)
	// number[:3] =  [0 1 2]
	// len=10 cap=10 slice=[0 1 2 3 4 5 6 7 8 9]

	fmt.Println("number[4:] = ", numbers[4:])
	printSlice(numbers)
	// number[4:] =  [4 5 6 7 8 9]
	// len=10 cap=10 slice=[0 1 2 3 4 5 6 7 8 9]

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	// len=0 cap=5 slice=[]

	numbers2 := numbers[:2]
	printSlice(numbers2)
	// len=2 cap=10 slice=[0 1]

	numbers3 := numbers[2:5]
	printSlice(numbers3)
	// len=3 cap=8 slice=[2 3 4]

	numbers4 := numbers[1:2]
	printSlice(numbers4)
	// len=1 cap=9 slice=[1]

	numbers5 := numbers[7:8]
	printSlice(numbers5)
	// len=1 cap=3 slice=[7]
	numbers6 := append(numbers5, 2, 3, 4, 5)
	printSlice(numbers6)
	// len=5 cap=6 slice=[7 2 3 4 5]
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
