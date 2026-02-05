package function

import (
	"errors"
	"fmt"
)

type Operations int
var myprovate = 3
const (
	SUM Operations = iota
	SUB
	MUL
	DIV
)


func Add(x int, z int) int {
	return x + z
}

func RepeatIncrement(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

func Calc(op Operations, x  , z float64) (float64, error){
	switch op {
	case SUM:
		return x + z, nil
	case SUB: 
		return x - z, nil	
	case MUL:
		return x * z, nil
	case DIV: 
		if z == 0{
			return 0, errors.New("no se puede dividir por  0")
		}
		return x / z, nil
	}
	return 0, errors.New("operacion no valida")
}

func Split(v int )(int, int ){
	y := v * 4	/ 2
	z := v * 12 - y
	return y, z
}

// otra forma de declarar funciones en variables 
func Split2 (v int )(x, y int ){
	x = v * 7/ 3
	y = v * 12 - 3
	return
}

func Msum (value ...float64)float64{
	var  sum float64
	for _, v := range value {
		sum += v
	}
	return sum
}

func MOperations(op Operations,values ...float64) (float64, error){
	if len(values) == 0 {
		return 0, errors.New("Solo valores correctos")
	}
	sum := values[0]

	for _, v := range values[1:]{
		switch op {
			case SUM:
				sum += v
			case SUB: 
				sum -= v
			case MUL:
				sum *= v	
			case DIV: 
				if v == 0 {
					return 0, errors.New("no se puede dividir por 0")
				}
			     sum /= v
		}
	}
	return sum, nil
}
