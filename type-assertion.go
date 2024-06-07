package main

import "fmt"

/*
	Type assetion adalah kemampuan untuk merubah tipe data apapun menjadi tipe data yg diinginkan
	Fitur ini sering digunakan ketika bertemu dengan data interface kosong
*/

func random() interface{} {
	return 10
}

func main(){
	// var result interface{} = random()
	// var resultString string= result.(string)
	// fmt.Println(resultString)

	/*
		Untuk lebih aman menggunakan tipe assertions, sebaiknya kita menggunkan switch experion
		Karena jika salah menggunakannya maka akan berakibat terjadi panic di apk kita dan tidak
		ter-recover(utk menangkap data panic) maka program akan mati

	*/
	var result interface{} = random()

	switch value := result.(type){
	case string :
		fmt.Println("Value", value, "is string")
	case int :
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Unknown Type")
	}
}
