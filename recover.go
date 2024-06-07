package main

import "fmt"
/*
	recover adalah func yg digunakan utk menangkap data panic/error
	dan menghentikan proses fun panic berhenti dan program akan tetap berjalan
*/

func endApp(){
	pesan := recover() //menangkap panic erornya
	
	if pesan != nil{
		fmt.Println("Terjadi Error", pesan)
	}
	fmt.Println("Apk berhenti")
}

func runApp( error bool){
	defer endApp()//dilewati sebelum code di bawanya dieksekusi (skip kode)
	if error {
		panic("ERROR!!!")
	}
	fmt.Println("Apk Berjalan")


}

func main(){
	runApp(true)
}