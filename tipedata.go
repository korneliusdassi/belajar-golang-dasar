package main


import "fmt"

func main(){
	// //tipe data : numerik/angka, string, float, const
	// umur := 25 //otomatis terbaca integer
	
	// fmt.Println("Umur saya", umur,"tahun")

	x:=0

	for i :=0; i<2; i++{
		for j :=0; j<2; j++{
			x+=i+j
		}
	}
	fmt.Println(x)
}