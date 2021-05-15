package main

import (
	"fmt"
	"time"
)

func main() {

	startTime1 := time.Now()
	//Complexity O(n)
	println(soma1(10))

	fmt.Println(time.Since(startTime1).Seconds(), "Time execution for{} method")

	startTime2 := time.Now()

	//Complexity O(3)
	println(soma2(10))

	fmt.Println(time.Since(startTime2).Seconds(), "Time execution for{} method")
}

func soma1(n int) int {
	soma := 0

	for i := 0; i <= n; i++ {
		soma += i
	}

	return soma
}

func soma2(n int) int {
	return (n * (n + 1)) / 2
}
