package main

import "fmt"

/*
	● Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam
		method adalah pass by value
	● Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena
		harus selalu diduplikasi ketika memanggil method
*/

type Man struct {
	Name string
}

func (man *Man) Status() {
	man.Name = "Mr. " + man.Name
}

func main() {
	lius := Man{"Kornelius"}
	lius.Status()
	fmt.Println(lius.Name)
}
