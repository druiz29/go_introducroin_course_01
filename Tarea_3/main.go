package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number:")
	var num int
	var arr []int

	fmt.Println("ingrese 0 para terminar")

	for {
		fmt.Scan(&num) 
			if num == 0 {
				break
			}
			arr = append(arr, num)
		}
		fmt.Println("Array before modification:", arr)
	}
