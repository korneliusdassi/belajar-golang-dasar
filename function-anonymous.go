package main

import "fmt"

type Blacklist func(string) bool //type declarations

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	//memasukkan anonymous function  kedalam variabel
	blackList := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blackList)
	registerUser("Dassi", blackList)

	//memaksukan anonymous function kedalam parameter secara langsung
	registerUser("Anjay", func(name string) bool {
		return name == "Anjay"
	})
}
