package main

import "fmt"

/*
	Pointer adalah kemampuan membuat reference kelokasi data di memory yang sama, tanpa
	menduplikasi data yang sama
	sederhananya kita bisa membuat pass by reference

	Untuk membuat sebuah variabel dengan nilai pointer ke variabel yang lain, kita bisa menggunakan
	operator/simbol & diikuti dengan nama variabel lainnya 
*/

type Address struct {
	City, Province, Country string
}

func main(){
	//contoh pointer

	address1 := Address{"Makasssar", "Sulsel", "Indonesia"}
	address2 := &address1 //mereferen ke address1(alamat memorinya)

	address2.City = "Toraja"

	*address2 = Address{"Rantepao", "Sulsel", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	/*
		Kita juga bisa membuat pointer menggunakan fungsi new selain menggunakan operator &
		Fungis new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/
	var address3 *Address= new(Address)
	address3.City = "Makale"
	fmt.Println(address3)

	fmt.Println("=========================================")

	//contoh lain
	var value = 100

	var valueAddres *int = &value //cara mengambil pointer menggunakan tanda * dan & untuk mereferensikannya

	//untuk mendapatkan alamatnya
	fmt.Println(valueAddres) //outputnya 0xc000018088 adl alamat dari memori alamat variabel itu disimpan
	//untuk mendapatkan nilainya kita menggunakan tanda *
	fmt.Println(*valueAddres)
	//untuk mendapatkan alamat var value cukup menggunkan tanda &
	fmt.Println(&value) // & artinya kita mereferens ke alamat memorinya
	
}