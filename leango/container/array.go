package container

import "fmt"

func printArray(arr *[5]int) {
	for i, v := range arr {
		//i为下标 v为值
		fmt.Println(i, v)
	}
}

//func main() {
//	var arr1 [5]int
//	arr2 := [3]int{1, 3, 5}
//	arr3 := [...]int{2, 4, 6, 8, 10}
//	var grid [4][5]int
//
//	printArray(&arr1)
//	//printArray(arr2) 报错
//	printArray(&arr3)
//}
