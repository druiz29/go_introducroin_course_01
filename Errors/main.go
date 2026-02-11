package main

import (
	"errors"
	"fmt"
	operator "go_introducroin_course_01/Errors/operators"
)

func main() {
	var err error
	err = errors.New("My err error")
	fmt.Println(err)

	err2 := fmt.Errorf("my format err, string:  %s, number: %d", "String f", 32)
	fmt.Println(err2)
	fmt.Println(err2.Error())
	fmt.Println()


	defer func (){
		r := recover()
		if r != nil {
			fmt.Println("Recovered in defer func")
			fmt.Println( "recovered in:" , r)
		}
	}()

	z := operator.Div(23, 3)
	fmt.Println("Result z is:",z)

}