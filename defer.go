package main

import "fmt"

func loggin(){
	fmt.Println("Slesai saat memanggil func")
}

func runApp(){
	defer loggin()//dilewati sebelum code di bawanya dieksekusi (skip kode)
	fmt.Println("Running apk")//di eksekusi dulu baru jalankan defer
}

func main(){
	runApp()
}