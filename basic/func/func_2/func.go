package main

import (
	"fmt"
	"math"
)

func main() {
	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(2))
}
