package main

import (
	"fmt"
	"go_introducroin_course_01/Functions/function"
)

func main() {
	fmt.Println( function.Add(2, 9))
	function.RepeatIncrement(3, "Hola Mundo")
	fmt.Println()
	v, err := function.Calc(function.SUM, 52,96)
	//fmt.Println("value :",r, "err: ", err)
	if  err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("value: " ,v)
	}
	
	m, err :=  function.Calc(function.MUL, 76, 90)
		if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("value: ", m)
	}
	
	d, err := function.Calc(function.DIV, 100, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("value: ", d)
	}

	y, z := function.Split(25)
	fmt.Println("Split: ", y, "value 2: ", z)


	v = (function.Msum(1,2,3))
	fmt.Println("multy sum:", v)

	v, err = function.MOperations(function.SUM, 1,6,7)
	fmt.Println("multy sum ", v, "err :", err)

	v , err = function.MOperations(function.SUB, 87,12,8)
	fmt.Println("multy sub ", v ,"err ", err)

	v, err = function.MOperations(function.MUL, 78,2,0)
	fmt.Println("multy mul:", v, "err", err)

	v, err = function.MOperations(function.DIV, 21,2)
	fmt.Println("multy div : ",v , "err", v)

	fn := function.FactoryOperation(function.SUB)
	v = fn(90, 45)
	fmt.Println("Factory Operation SUB :", v)

	fn = function.FactoryOperation(function.MUL)
	v = fn(100, 45)
	fmt.Println("Factory Operation MUL :", v)

}
