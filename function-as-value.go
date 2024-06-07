package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	//merubah fungsi getGoodBye menjadi sebuah value/nilai
	goodbye := getGoodBye
	// fmt.Println(goodbye("Kornel"))
	result := goodbye("Kornel")
	fmt.Println(result)
}
