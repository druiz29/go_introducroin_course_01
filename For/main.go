package main

import (
	"fmt"
)
func main (){
	sum := 0
	
	for  i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)

	sum = 1

	for sum < 100 {
		sum ++
	}
	fmt.Println(sum)

	sum = 0 
	for {
		if sum > 999 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	arr := []int {1, 2, 3, 4,}

	for i  := range arr {
		fmt.Println("index is : ", i, "value is : ", arr[i])
	}

	for i, v := range arr {
		fmt.Println("value is : ", v, "index is: ",i)
	}

	for _, v := range arr {
		fmt.Println("value is : ", v)
	}

	map1 := map[string]float64 {
		"A" :12.3,
		"B" : 1,
		"C" : 21.2,
	}
	 
	for k, v :=  range map1 {
		fmt.Println("key: ",k, "value: ", v)
	}

	for  i := range map1 {
		fmt.Println("key: ", i)
	}

	map3 :=  map[string][]int {
		"X": nil,
		"Y": { 1, 20, 4, 5 },
		"Z": { 54, 6, 66, 34, 4343 }, 
	}

	for k, v  := range map3 {
		fmt.Println("Key:",  k)
		for _, v := range v {
			fmt.Println("Value:", v)
		}
		fmt.Println()
	}


	
}