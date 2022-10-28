package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello World")

    // deklarasi variabel
	nama:="chintya"
	var angkaInt int = 10
	var decimalNumber = 36.6689898928

    // printf dipake buat formatting
	fmt.Printf("Decimal number %f\nNama saya %s\n Dan umurku %d \n\n", decimalNumber,nama,angkaInt)

    // printf dipake buat formatting
	fmt.Printf("Decimal number 4 angka dibelakang koma %.4f\n\n",decimalNumber)

	// bool
	var exist bool = false
	fmt.Printf("exist? %t\n\n",exist)

	// pake backtics
	message:=`Nama saya "chintya"
	umur saya 21 tahun
	saya suka kamu`

	fmt.Println(message)

	// kalau variabel di dekalrasi 2 kali , yang keluar pasti deklarasi terakhir 

	nama ="chintya"
	nama ="suciw"
	// kalo semisal nama=89 bakal error karena initial valuenya string tapi distu diganti integer
	fmt.Println("\nNama saya "+nama+"\n")


	// deklarasi multivariable
	var namaDepan, namaBelakang string ="Chintya","Suci"

	umurSaya, umurKamu := 21 , 12

	fmt.Printf("Nama saya %s %s \n Umur saya%d \n Umur kamu %d\n\n",namaDepan,namaBelakang,umurSaya,umurKamu)

	// deklaration underscore atau trash variabel digunakan untuk buang nilai
	// digunakan kalo kita ambil var yang dari db bukan di script ini , 
	// soalnya golang itu strict buat kayak gini 
	// kalo di declare 2 nnti error
	// dia dialokasi memori juga ilang 
	_=namaBelakang
	_="golang itu mudah"
	namaNama,_ := "Chintya","Suci"

	fmt.Printf("Nama saya %s \n\n",namaNama)

	// var bisa diubah nilainya tapi kalu const engga 
	// const bisa digunakan untuk inisial variable
	const umurConst int = 20
	// umurConst=46
	fmt.Printf("Umur saya : %d\n\n",umurConst)

	// conditional 
	// switch case

	

	// looping 
	var fruits = [4]string{"apel","melon","anggur","tomat"}
	// opsi 1 kalo ada index nya
	for i, fruit := range fruits {
		fmt.Printf("elemen ke %d = %s \n\n", i, fruit)
	}
	// opsi kalo gaada index [1]
	// inget !! go itu strict sama deklarasi variable gbisa di def tapi g dipake
	for _, fruit := range fruits {
		fmt.Printf("elemen %s \n\n", fruit)
	}
	
	// opsi kalo gaada index [2]
	// for fruit := range fruits {
	// 	fmt.Printf("elemen %s \n\n", fruit)
	// }

	// function di go ada definisi keluaran / return 



	//go ga strict sama 
	
	// struck digunakan untuk nyimpen data yang jenisnya bisa beda-beda

	
	// seperti array tapi index buat ngaksesnya bisa di define sendiri 
	// contoh dibawah ini dia string index nya
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] =  90
	chicken["februari"] = 70

	fmt.Println("januari : ", chicken["januari"])
	fmt.Println("februari : ", chicken["februari"])

	var names [4]string 
	names[0]="bro"
	names[1]="bri"
	names[2]="bre"
	names[3]="sis"
	

	fmt.Printf(names[1],names[3])
	fmt.Println(names)


	fmt.Println("fruits[0.2] =>", fruits[0:2])
}