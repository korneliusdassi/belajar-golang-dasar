package main


import (
	"fmt"
)

//1.deklarasi variabel
// var namaDepan = "Kornelius"
// var namaBelakang = "Dassi"

//2.deklarasi variabel
// namaDepan := "Kornelius"->hanya boleh dipakai didalam blok fungsi
var namaDepan, namaBelakang = "Kornelius","Dassi"

//3. deklarasi variabel tanpa nilai
// var nama string

func main(){
	fmt.Println("Perkenalkan Nama Lengkap Saya "+ namaDepan+" "+namaBelakang)
}
