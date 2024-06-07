package main

import "fmt"

func main(){
	/*
		nilai := 90

		if nilai <= 100 {
			fmt.Println("Selamat anda Lulus dengan nilai: ",nilai)
		}else{
			fmt.Println("Maaf anda belum bisa dinyatakan lulus")
		}

	*/

	/*
		membandingkan nilai tipe string
	
		warna := "biru"
		if warna == "biru"{
			fmt.Println("Anda Benar")
		}else{
			fmt.Println("Anda Salah")
		}
	*/

	/*kondisi jika bernilai boolean
		if true {
			fmt.Println("lanjut")
		}else{
			fmt.Println("tidak bisa lanjut")
		}
	*/

	/*
		if dengan short statement mirip dengan for loop
	*/
	nama := "Kornelius"
	if panjang := len(nama); panjang > 10 {
		fmt.Println("Nama terlalu panjang!")
	} else{
		fmt.Println("Nama sudah benar!")
	}


}