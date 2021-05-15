package operations

import (
	"fmt"
	"reflect"
	"sort"
)

func ArrayInt(slice []int) {

	fmt.Println("Original array:", slice)
	sort.Ints(slice)
	fmt.Println("Sorted array:", slice)

	fmt.Println("Print up to 5 element:", slice[:5])
	fmt.Println("Type of array: ", reflect.ValueOf(slice).Kind())

	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(slice); i++ {
		fmt.Print(" ", slice[i])
	}

	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range slice {
		fmt.Println(index, "=>", element)
	}

	fmt.Println("\n---------------Example 3--------------------\n")
	for _, value := range slice {
		fmt.Print(" ", value)
	}

	fmt.Println()

	fmt.Println("\n---------------Example 4 Remove--------------------\n")
	for index, element := range slice {
		if element == 5 {
			slice = append(slice[:index], slice[index+1:]...)
			fmt.Println("Array without element 5:", slice)
		}
	}

	fmt.Println("\n---------------Example 5 Major Element in the Array--------------------\n")
	fmt.Println("The Major Element is: ", MajorValueOnArray(slice))
}

func ArrayString(slice []string) {

	fmt.Println("Original array:", slice)
	sort.Strings(slice)
	fmt.Println("Sorted array:", slice)

	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(slice); i++ {
		fmt.Print(" ", slice[i])
	}

	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range slice {
		fmt.Print(" ", index, " => ", element, " |")
	}

	fmt.Println("\n---------------Example 3--------------------\n")
	for _, element := range slice {
		fmt.Print(" ", element)
	}

	fmt.Println("\n---------------Example 4 Remove--------------------\n")
	for index, element := range slice {

		if element == "D" {
			fmt.Println("Element in the index", index, " was removed: ", RemoveValueArray(slice, index))
		}

	}

}

func MajorValueOnArray(a []int) int {
	mValue := 0

	for _, element := range a {
		if element > mValue {
			mValue = element
		}
	}

	return mValue
}

func RemoveValueArray(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
