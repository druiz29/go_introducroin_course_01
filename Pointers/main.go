package main

import (
	"fmt"
)

type myStruct struct {
	Id int 
	Name string
}
func (m myStruct) Set(name string){
	fmt.Printf("addrs: %p\n", &m)
	m.Name = name
}

func (m *myStruct) SetPointer(name string){
	fmt.Printf("addrs: %p\n", m)
	m.Name = name
}

func main() {
	var i  int
	var iP *int
	i = 12
	
	iP = &i
	
	fmt.Println(&i)
	fmt.Println(iP)
	fmt.Println(i)
	fmt.Println(*iP)

	*iP = 21

	fmt.Printf("val i : %d, val pointer i : %d\n", i, *iP)

	myVar := 40
	fmt.Printf("addrs: %p, values: %v\n", &myVar, myVar)
	increment(myVar) 	 	
	fmt.Println(myVar)
	incrementP(&myVar)
	fmt.Println(myVar)     
	fmt.Println()      
	
	var mySlice []int 
	mySlice = append(mySlice, 3,4,2)
	fmt.Printf("addrs: %p, values: %v\n", &mySlice, mySlice)
	fmt.Printf("addrs 1: %p, addrs 2: %p, addrs 3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println("mySlice:", mySlice)
	fmt.Println()      
	myStruct := &myStruct{ Id: 1234, Name: "test"}
	fmt.Printf("addrs: %p\n ",myStruct)
	fmt.Printf("id: %d name: %s\n", myStruct.Id, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("id: %d name: %s\n", myStruct.Id, myStruct.Name)
	fmt.Println()      
	fmt.Printf("addrs: %p\n", myStruct)
	myStruct.Set("test method 1")
	fmt.Printf("id: %d name: %s\n", myStruct.Id, myStruct.Name)
	myStruct.SetPointer("test method pointer 2")
	fmt.Printf("id: %d name: %s\n", myStruct.Id, myStruct.Name)

}


func increment(val int) {
		fmt.Println(&val)
		val++
	}

func incrementP(val *int ) {
		fmt.Println(val)
		*val++
	}
func updateSlice( v []int){
	fmt.Printf("addrs: %p\n", v)
	fmt.Printf("addrs: 1 %p, addres 2: %p, addres 3: %p", &v[0], &v[1], &v[2])
	v[0] = 12
}

func updateStruct(s *myStruct){
	fmt.Printf("addrs struct: %p\n", s)
	s.Id = 123
	s.Name = "updated"
}
