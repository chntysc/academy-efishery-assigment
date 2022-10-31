package main

import (
	"fmt"
)

type student struct {
	name string
	grade int
}


func main(){

	// var input =10

	var numberA int = 4
	var numberB *int = &numberA
	
	
	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	
	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	



	var s1 = student{name: "wick", grade: 2}
	
	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
	
	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)

}

