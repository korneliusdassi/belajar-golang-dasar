package main

import "fmt"

func main() {
	days := [...]string{ //array
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	slice := days[3:6] //mengambil slice dari array

	fmt.Println(slice)
	fmt.Println("Panjang slice =", len(slice))
	fmt.Println("Kapasitas slice =", cap(slice))

	//cara membuat slice paling aman dri awal krn si arraynya ada di dlm slicenya dan gk keliatan
	name := make([]string, 2, 5) //artinya panjang arraynya 5

	name[0] = "Kornelius"
	name[1] = "dassi"
	fmt.Println(name)

	//cara meng copy slice
	//copy(toSLice, fromSlice
	copySlice := make([]string, len(name), cap(name))
	copy(copySlice, name)

	fmt.Println(copySlice)
}
