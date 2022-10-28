package main

import (
	"fmt"
)


func main(){

	// var input =10

	var input int
	fmt.Println("Masukkan angka : ")
	fmt.Scan(&input)

	if input%2==0{
		fmt.Println("angka yang dimasukkan genap \n")
		genap(input)
	} else {
		fmt.Println("angka yang dimasukkan ganjil \n")
		ganjil(input)
	}

}

func ganjil(num int){
	var a =""
	for i := 0; i < num; i++ {
		a = a + "*"
		fmt.Println(a + "\n")
	}
}

func genap(num int){
	var x =""
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			x = x + "*"
		}
		fmt.Println(x + "\n")
		x =""
	}
}