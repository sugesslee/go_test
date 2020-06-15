package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/09/10     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/9/10 4:06 PM
 * @date 2019/9/10 4:06 PM
 * @since 1.0.0
 */

func eval1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div1(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unSupported operation: %s", op)
	}
}

func div1(a, b int) (q, r int) {
	return a / b, a % b
}

func sum1(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func apply1(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("calling function %s with args (%d, %d)\n", opName, a, b)

	return op(a, b)
}

func main() {
	q, r := div1(13, 3)
	fmt.Printf("q / r is %d ... %d\n", q, r)

	if result, err := eval1(3, 4, "-"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

	fmt.Printf("numbers sum is: %d\n", sum1(1, 2, 3, 4, 5))

	fmt.Println("pow(3, 4) is: ", apply1(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
}
