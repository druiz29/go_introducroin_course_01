package main

import (
	"fmt"
)
func main() {

	m1 := make(map[int]string)
	m1[1] = "A"
	m1[2] = "B"
	m1[3] = "C"

	fmt.Println(m1)
	fmt.Println(m1[2])

	delete(m1,3)
	fmt.Println(m1)

	m1[1] = "D"
	fmt.Println(m1[1])

	m1[7] = " "
	fmt.Println(m1[7])
	fmt.Println(m1[33])

	v , ok := m1[7]
	fmt.Println("value is : ", v, "ok is :", ok)

	v, ok = m1[33]
    fmt.Println("value is : ", v, "ok is :", ok)

	m2 := map[int]string {
		6: "F",
		8: "G",
		9: "H",
	}

	fmt.Println(m2)
}
