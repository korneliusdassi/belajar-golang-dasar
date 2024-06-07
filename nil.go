package main

import "fmt"
/*
	Nil hanya boleh digunakan ditipe data seperti map, slice, function, interface, pointer, dan channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}else{
		return map[string] string{
			"name" : name,
		}
	}
}

func main(){
	// person := NewMap("Lius")
	var person map[string] string = NewMap("Lius")

	if person == nil {
		fmt.Println("Data Kosong")
	}else{
		fmt.Println(person)
	}
}