package main

import "fmt"

/*
	function recursive adalah fungsi yang memanggil fungsi dirinya sendiri
	atau memanggil fungsi dalam fungsinya sendiri
*/

//tanpa fungsi rekursif
func factorial(value int) int {
	result := 1
	for i := value; i > 0; i--{
		result *= i
	}
	return result
}

//menggunakan fungsi rekursif
func factorialRecursive(nilai int) int {
	if nilai == 1 {
		return 1
	}else{
		return nilai * factorialRecursive(nilai-1)
	}
}

func main(){
	loop := factorial(5)
	fmt.Println(loop)
	fmt.Println("------------------------------")
	recursive := factorialRecursive(5)
	fmt.Println(recursive)

}