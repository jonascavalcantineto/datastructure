package rotate

import "fmt"

func RotLeft(a []int32, d int32) []int32 {

	aTotal := len(a)

	for i := int32(0); i < d; i++ {
		aNew := make([]int32, len(a))

		aNew[aTotal-1] = a[0]

		for index := 1; index < aTotal; index++ {
			aNew[index-1] = a[index]
		}
		a = aNew

		aNew = nil
		fmt.Println(a)
	}

	return a
}
