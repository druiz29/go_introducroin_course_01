package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ingrese Codigo Valido (10-15-21-27-35):")
	var array []string
	fmt.Println("Ingrese 0 para terminar")

	for {
			var value string
			fmt.Scan(&value) 
			
			if value == "0" {
				break
			}

			if value ==  "10"  {
				array = append(array, "notebook")
			} else if value == "15" {
				array = append(array, "tv")
			} else if value == "21" {
				array = append(array, "heladeria")
			} else if value == "27" {
				array = append(array, "monitor")
			} else if value == "35" {
				array = append(array, "camara")
			} else   {
				array = append(array, "No encontrado")
			}


			// switch value {
			
			// case "10":
			// 	array = append(array, "notebook" )
			// case "15":
			// 	array = append(array, "tv" )
			// case "21":
			// 	array = append(array, "heladeria" )
			// case "27" :
			// 	array = append(array, "monitor")
			// case "35": 
			//     array = append(array, "camara")		
			// default:
			// 	array = append(array, "No encontrado")	
			// }
			fmt.Println("Array codes:", array)
		}
}