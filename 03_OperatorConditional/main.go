package main

import "fmt"
func main (){
	years_old := 18
	
	fmt.Println("Operators")
	fmt.Println(years_old > 17) // true
	fmt.Println(years_old < 18) // false 
	fmt.Println(years_old <= 18) // true 
	fmt.Println(years_old >= 20) // false
	fmt.Println(years_old == 18) // true 

	fmt.Println("]OR")
	// OR
	fmt.Println(years_old < 18 || years_old == 18) // true
	fmt.Println(years_old < 18 || years_old == 19) // true
	fmt.Println(years_old < 20 || years_old == 18) // true
	
	fmt.Println("&&")
	// &&

	fmt.Println(years_old < 18 && years_old == 18) // 
	fmt.Println(years_old < 18 && years_old == 19) //
	fmt.Println(years_old < 20 && years_old == 18) // 

	fmt.Println("!")
	// !
	fmt.Println(true)
	fmt.Println()

}