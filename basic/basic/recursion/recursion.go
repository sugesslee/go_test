package main

import "fmt"

func main() {
	fmt.Println("5! = ", Factorial(5))
	fmt.Println("fib 5 = ", fibonacci(5))
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n - 2 + fibonacci(n-1))
}
