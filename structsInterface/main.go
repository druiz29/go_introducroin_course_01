package main

import (
	"encoding/json"
	"fmt"
	"go_introducroin_course_01/structsInterface/structs"
	"go_introducroin_course_01/structsInterface/vehicles"
)

func main() {
	var p1 structs.Product
	p1.ID = 1
	p1.Name = "Termo"
	fmt.Println(p1)

	p2 := structs.Product{}
	p2.ID = 2
	p2.Name = "vaso"
	fmt.Println(p2)

	p3 := structs.Product{3,"jabon",structs.Type{},5000,5, nil}
	fmt.Println(p3)

	// buena practica 

	p4 := structs.Product {
		ID:  4,
		Name: "Soda",
		Price: 3000,
	}
	fmt.Println(p4)

	p5 := structs.Product {
		Name: "Pruductos HyG",
		Type: structs.Type{
			Code: "A",
			Description: "Aseo",
		},
		Tags: []string{"Aseo", "Despensa","licores","Verduras"},
	}
	fmt.Println(p5)

	P6 := structs.Product {
		// ID : 4,
		Name: "Pruductos HyG",
		Price: 21434,
		Type: structs.Type{
			Code: "A",
			Description: "Aseo",
		},
		Tags: []string{"Aseo", "Despensa","licores","Verduras"},
		Count: 2,
	}
	v, err := json.Marshal(P6)
	fmt.Println(err)
	fmt.Println(string(v))

	fmt.Println("Total Price", P6.TotalPrice())
	fmt.Println(P6)
	P6.SetName("other name:")
	P6.AddTags("tag1", "tag2", "tag3")
	fmt.Println(P6)

		P7 := structs.Product {
		// ID : 4,
		Name: "Pruductos JUD",
		Price: 1500,
		Type: structs.Type{
			Code: "A",
			Description: "Aseo",
		},
		Tags: []string{"Aseo", "Despensa","licores","Verduras"},
		Count: 22,
	}

		P8 := structs.Product {
		// ID : 4,
		Name: "Pruductos PERA",
		Price: 2000,
		Type: structs.Type{
			Code: "A",
			Description: "Aseo",
		},
		Tags: []string{"Aseo", "Despensa","licores","Verduras"},
		Count: 34,
	}

	car := structs.NewCar(00114)
	car.AddProducts(P6,P7,P8)
	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total products: ", len(car.Products))
	fmt.Printf("total car: $%.2f\n", car.Total())

	fmt.Println()
	fmt.Println("----- VEHICLES -----")


	carV := vehicles.Car{Time: 120}
	fmt.Println(carV.Distance())
	
	Varray := []string {"CAR", "MOTORCYCLE","TRUCK","AIRPLANE", "GOKU"}
	
	var d float64

	for _, v := range Varray {
		fmt.Printf("vehicle:  %s\n",v)
		vech, err := vehicles.New(v, 100)
		if err != nil {
		 	fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}
		distance := vech.Distance()
		fmt.Printf("Distance: %.2f\n",distance)
		fmt.Println()

		d += distance
	}
	fmt.Printf("Total distance: %.2f\n", d)
	fmt.Println()


}