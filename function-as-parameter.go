package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string){
// 	nameFilter := filter(name)
// 	fmt.Println("Hello", nameFilter)
// } atau

/*
	//type declarations
	type filter string // artinya filter adalah alias dari tipe data string
*/
type Filter func(string) string //rekomendasi utk func as parameter

func sayHelloWithFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("Hello", nameFilter)
}


func spamFilter(name string) string{
	if name == "Asuu" {
		return "..."
	}else{
		return name
	}
}

func main(){
	sayHelloWithFilter("Asuu", spamFilter)
	sayHelloWithFilter("Kornelius", spamFilter)
}