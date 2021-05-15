package main

import (
	"../datastructure/array/operations"
	"../datastructure/array/rotate"
	"fmt"
)

func main() {

	slice := []int{5, 2, 1, 6, 9, 8, 7, 3, 4}
	fmt.Println("Slice : ", slice)
	operations.ArrayInt(slice)

	slice32 := []int32{5, 2, 1, 6, 9, 8, 7, 3, 4}
	fmt.Println(rotate.RotLeft(slice32, 4))
}
