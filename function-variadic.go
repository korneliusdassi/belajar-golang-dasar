package main

import "fmt"

func sumAll(numbers ...int) int{
	total := 0
	for _, value := range numbers{
		total += value
	}
	return total
}

func main(){
	total := sumAll(10,20,30,40,50,60)
	fmt.Println(total)

	/*
		kadang ada kasus dimana kita menggunakan fungsi variadic, 
		namun memiliki variabel berupa slice
		kita bisa memasukkan slice ke dalam fungsi variadic
		yang di namakan slice parameter
	*/
	slice := []int{20,40,60}
	total = sumAll(slice...)
	fmt.Println(total)

}