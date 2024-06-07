package main

import "fmt"

func endApp(){
	fmt.Println("Apk berhenti")
}

func runApp( error bool){
	defer endApp()//dilewati sebelum code di bawanya dieksekusi (skip kode)
	if error {
		panic("error")//digunakanan utk menghentikan program dan biasa digunakan ketika terjadi eror saat program berjalan...saat func panic dipanggil program akan henti tapi func defer akan tetap jalan
	}
	fmt.Println("Apk Berjalan")
}

func main(){
	runApp(false)
}