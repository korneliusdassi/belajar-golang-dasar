package main

import "fmt"

func main() {
	/*
		map sangat mirip dengan object yaitu menggunakan format key value.
		yang membedakan adalah Map memperbolehkan key dengan tipe data apa pun,
		dibandingkan Object yang hanya mengizinkan key bertipe String atau Symbol
	*/

	var member = map[int]string{ //key[int] value[string]
		1: "Kornelius",
		2: "Dassi",
		3: "Kornel",
		4: "Lius",
	}
	/*
		menambahkan data ke dalam map
	*/
	member[5] = "Acci"

	fmt.Println(member)      //mengambil semua data beserta key nya
	fmt.Println(member[3])   //mengambil data tertentu dgn key
	fmt.Println(len(member)) //mengambil panjang map

	fmt.Println("---------------------------------")

	/*
		mengambil data key dan value satu2 menggunakan for loop
	*/
	for id, nama := range member {
		fmt.Println(id, nama)
	}

	fmt.Println("---------------------------------")

	/*
		mengambil data value nya saja tanpa mengikutkan key nya
	*/
	for _, nama := range member {
		fmt.Println(nama)
	}

	fmt.Println("---------------------------------")

	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]
	if isExist {
	fmt.Println(value)
	} else {
	fmt.Println("item is not exists")
	}

}
