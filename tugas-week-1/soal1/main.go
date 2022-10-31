package main

import (
	"fmt"
	"sort"
)


func main(){

	// database barang
	type produkEfishery struct {
		id int
		barang string
		harga  int
	}

	var dataBarangEfishery = []produkEfishery{
		{id: 1, barang : "Benih Lele", harga: 50000},
		{id: 2, barang : "Pakan lele cap menara", harga: 25000},
		{id: 3, barang : "Probiotik A", harga: 75000},
		{id: 4, barang : "Probiotin Nila B", harga: 10000},
		{id: 5, barang : "Pakan Nila", harga: 20000},
		{id: 6, barang : "Benih Nila", harga: 20000},
		{id: 7, barang : "Cupang", harga: 5000},
		{id: 8, barang : "Benih Nila", harga: 30000},
		{id: 9, barang : "Benih Cupang", harga: 10000},
		{id: 10, barang : "Probiotin B", harga: 10000},
	}

	sort.SliceStable(dataBarangEfishery, func(i, j int) bool {
		return dataBarangEfishery[i].harga < dataBarangEfishery[j].harga
	})

	var total = 0
	for _, barangEfishery := range dataBarangEfishery {
		switch {
		case total <= 100000 :
			total = total + barangEfishery.harga
			if total <=100000 {
				fmt.Println(barangEfishery.barang ," - " , barangEfishery.harga , "\n")
			} else {
				total = total - barangEfishery.harga
			}
		case total > 100000 :
			break
		}

	}
	fmt.Println("Total harga produk yang dibeli adalah " ,  total , "\n")


	fmt.Println("------------------------------------------------------"+ "\n")
	// daftar produk harga 10.000
	fmt.Println("Daftar produk dengan harga Rp 10000"+ "\n")
	for _, barangEfishery := range dataBarangEfishery {
		switch {
		case barangEfishery.harga == 10000 :
			fmt.Println(barangEfishery.barang , " - " ,barangEfishery.harga , "\n")
		}
	}

	fmt.Println("------------------------------------------------------"+ "\n")
	// min max
	var min, max int
	var varMin, varMax string
	for i, barangEfishery := range dataBarangEfishery {
		switch{
		case i == 0:
			max, min = barangEfishery.harga,barangEfishery.harga
		case barangEfishery.harga > max :
			max = barangEfishery.harga
			varMax = barangEfishery.barang
		case barangEfishery.harga <min :
			min = barangEfishery.harga
			varMin = barangEfishery.barang
		}
	}
	fmt.Println("Daftar produk termurah : " ,varMin , " Rp " ,min , "\n")
	fmt.Println("Daftar produk termahal : " , varMax , " Rp " ,max , "\n")

	

}



