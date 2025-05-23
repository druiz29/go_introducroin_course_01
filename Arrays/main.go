package main

import (
	"fmt"
	"unsafe"
)
func main() {
	var varInt int
	fmt.Println(varInt)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", varInt, unsafe.Sizeof(varInt),unsafe.Sizeof(varInt)*8 )

	var myArray [6]int
	fmt.Println(myArray)

	var myArray2 = [3]string{"julian","diego","oscar"}
	fmt.Println(myArray2)
	
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 5
	fmt.Println(myArray)
	fmt.Println("size:",len(myArray))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArray, unsafe.Sizeof(myArray), unsafe.Sizeof(myArray)*8)


	var slice1  []int 
	fmt.Printf("size: %d\n, value: %v\n", len(slice1),slice1)

	slice1 = append(slice1, 10,20,30,40)
	fmt.Printf("size: %d\n, value: %v\n", len(slice1),slice1)

	mySlice := myArray[1:3]
	fmt.Println(mySlice)
	fmt.Printf("size: %d\n, value: %v\n", len(mySlice),mySlice)
	fmt.Println("size :", len(mySlice))

	fmt.Println(&myArray[2])
	fmt.Println(&mySlice[0])

	fmt.Println(myArray)

	mySlice[0] = 80
	mySlice[1] = 100
	fmt.Println(mySlice)

	fmt.Println(myArray[:4])
	fmt.Println(myArray[2:])

	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Printf("size: %d, value: %v\n", len(slice),slice)
}