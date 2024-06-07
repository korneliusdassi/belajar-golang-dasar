package main

import "fmt"

func main(){
	// var n int
	// fmt.Print("Input Nilai: ")
	// fmt.Scan(&n)
	// fmt.Println()

	// for i := 1; i <= n; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := n; k >= i; k-- {
	// 		fmt.Print("*")
	// 	}
	// 	for m := n-1; m >= i; m-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	//implementasi for loop di array

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	for i:=0; i<len(days); i++{
		fmt.Println(days[i])
	}

	fmt.Println("====================================")

	//meluping nilai array menggunakan range
	for _, data := range days {
		fmt.Printf("Hari %s\n", data)
	}
}