package main

import (
	"fmt"
)
func main() {
	var age int = 15
	var license bool = false

	if license == true && age >= 15 {
		fmt.Println("Puede seguir avanzando")
	}else if age <= 15 || license == false {
		fmt.Println("No puede seguir avanzando")
	}
}