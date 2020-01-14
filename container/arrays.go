package main

import "fmt"

func arrayDefinition() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}
	var grad [4][5]bool
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grad)

	fmt.Println("printArr(arr1)")
	printArr(&arr1)
	fmt.Println("printArr(arr3)")
	printArr(&arr3)
	fmt.Println("arr1 + arr3")
	fmt.Println(arr1, arr3)
}

func printArr(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	arrayDefinition()
}
