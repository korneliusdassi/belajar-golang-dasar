package main

import "fmt"

/*
	1.interface adlh tipe data abstrak, dia tdk memiliki implementasi langsung
	2.sebuah interface berisikan defenisi2 atau kumpulan method/func
	3.biasanya interface digunkan sebagai kontrak
*/

type HasName interface{
	GetName() string //returnnya adl string
}
func SayHello(hasName HasName){
	 fmt.Println("Hello", hasName.GetName())
}

//implementasi 1
type Person struct{
	Name string
}

//dibawah ini adalah person yg punya kontrak si GetName
func (person Person) GetName() string{
	return person.Name
}

//implementasi 2
type Animal struct{
	Name string
}

func (animal Animal) GetName() string{
	return animal.Name
}

func main(){
	var lius Person
	lius.Name = "Kornelius"
	SayHello(lius)

	cat := Animal{
		Name : "Kitten",
	}
	SayHello(cat)
}