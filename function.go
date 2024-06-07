package main

import (
	"fmt"
	// "strconv"
)

/*
	//fungsi tanpa parameter
	func printMessage() int{
		return 10
	}

	func main(){
		fmt.Println("Fungsi menghasilkan nilai", printMessage())
	}

*/

/*
	//fungsi menggunakan parameter

	func biografi(nama string, status string, umur int) string {
		//convert int ke string
		newUmur := strconv.Itoa(umur)
		return nama +" adalah seorang "+ status + " dan saat ini berumur " + newUmur +" tahun"
	}

	func main(){
		fmt.Println("Perkenalkan",biografi("kornelius","mahasiswa",25))
	}
*/

/*
	//fungsi parameter dengan return banyak nilai(multiple return)

	func biografi(nama string, status string, umur int) (string, int){
		bio := nama +" "+ status
		return bio, umur
	}

	func main(){
		info, umur := biografi("Dassi","mahasiswa ",25)
		fmt.Println(info, umur)
	}
*/

/*
//fungsi tanpa return
var text = "ini adalah teks"
func printText(){
	fmt.Println(text)
}
func main(){
	printText()
}
*/

// fungsi return named value

func getFullName() (firstName, LastName string) {
	firstName = "Kornelius"
	LastName = "Dassi"

	return
}

func main() {
	namaDepan, namaBelakang := getFullName()
	fmt.Println(namaDepan, namaBelakang)
}
