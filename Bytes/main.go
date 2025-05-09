package main

import (
	"fmt"
)

func main () {
	var A  byte = 'A'
	var a byte = 'a'
	var R byte = 82
	var s byte = 115
	vector := []byte{ 65,97,82,115}  
	
	fmt.Println("value in code ASCII: ")
	fmt.Println(A)
	fmt.Println(a)
	fmt.Println(R)
	fmt.Println(s)

	fmt.Println()
	fmt.Println("value string in code ASCII: ")
	fmt.Println(string(A))
	fmt.Println(string(a))
	fmt.Println(string(R))
	fmt.Println(string(s))
	fmt.Println(string(vector))

	// valor de bytes

	fmt.Printf()

}