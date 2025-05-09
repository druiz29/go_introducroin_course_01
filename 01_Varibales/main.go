package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myIntVar  int 
	myIntVar = 12
	fmt.Println("My variable is: ",myIntVar)

	var myUnintVar uint 
	myUnintVar= 10
	fmt.Println("My variables is:", myUnintVar)

	var myStringVar string
	myStringVar = "Hello zzzzz"
	fmt.Println("my variables string is : ", myStringVar)

	var myBoolVar bool 
	myBoolVar = false
	fmt.Println("my variable boolean is: ", myBoolVar)

	fmt.Println("my variable boolean is: ", &myStringVar)

	myIntVar2 := 34
	fmt.Println("my variable not defined is: ", myIntVar2)
	
	myStringVar2 := "Hello World"
	fmt.Println("My varible not defined is: ", myStringVar2)


		// string es una secuencia de bytes
		var myStringVar3 string
		fmt.Printf("String default value: %s\n", myStringVar3)
	
		myStringVar5 := `My string variable in golang
		with multiple
		line`

		fmt.Printf("The variable value is: %s\n", myStringVar5)

			// con las llaves definimos un scope

	{
		//comversiones de varibales
		
		fmt.Println()
		
		// comversiones de float a string
		floatVar := 9.2339427
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.3f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar) 

		//comversion de enterp a string
		intVar := 45
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		instStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", instStrVar, instStrVar)

		// comversion de string a int 
		strVar, err := strconv.ParseInt("643", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", strVar, strVar)
		
		// comversion de string a float
		strVar2, _ := strconv.ParseFloat("-9754.352", 64) 
		fmt.Printf("type: %T, value: %.2f\n", strVar2, strVar2)

		// Bytes
		//  


	}
}
