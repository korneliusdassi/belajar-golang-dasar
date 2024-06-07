package main

import "fmt"

func Ups(i int) interface{} { //return valuenya adalah interfacenya kosong
	if i == 1{
		return 1
	}else if i == 2{
		return true
	}else{
		return "Ups"
	}
}

func main(){
	//harus set sbgi interface kosong krn return valuenya interface kosong
	var data interface{} = Ups(3)
	fmt.Println(data)
}