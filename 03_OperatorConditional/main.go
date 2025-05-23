package main

import "fmt"
func main (){
	years_old := 30
	
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
	fmt.Println(!true)
	
	fmt.Println(years_old < 40)
	fmt.Println(!(years_old > 40))

	fmt.Println(years_old < 25 && years_old == 32 || years_old < 40 ) // false
	fmt.Println(years_old < 25 && (years_old == 32 || years_old < 40)) // false	


	years_old = 20
	if years_old > 18 {
		fmt.Println("%d You are an adult\n", years_old)
	}
	boolVal := true

	if boolVal {
		fmt.Println("boolVal is true")
	} else {
		fmt.Println("boolVal is false")
	}

	if  value := true; value { 
		fmt.Println("value is true")
	}

	number := 0

	if number == 1{
		fmt.Println("number is 1")
	}else if number == 2 {
		fmt.Println("number is 2")
	}else if number == 3 {
		fmt.Println("number is 3")
	}
	
	switch number {
	case 1:
		fmt.Println("number is 1")
	case 2:
		fmt.Println("number is 2")
	case 3:
		fmt.Println("number is 3")
	default:
		fmt.Println("number is not 1, 2 or 3")		
	}

	switch {
		case number < 0:
			fmt.Println("number is negative") 
		case number > 0:
			fmt.Println("number is positive")
		case number == 0:
			fmt.Println("number is zero")		
	}


}