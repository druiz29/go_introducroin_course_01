package main

import "fmt"

func main() {

	arr := [5]int{4, 2, 5, 6, 7}

	for i := range arr {
		arr[i] += 20
	}

	fmt.Println(arr)



	arr2 := [5]int{4, 2, 5, 6, 7}
	for i := range arr2 {
		arr2[i] *= 2
	}

	fmt.Println(arr2)
}