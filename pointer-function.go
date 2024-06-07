package main
import "fmt"

/*
	● Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di
		copy lalu dikirim ke function tersebut
	● Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah
		berubah.
	● Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
	● Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
	● Untuk melakukan ini, kita juga bisa menggunakan pointer di function
	● Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di
	parameternya
*/

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address){
	address.Country= "Indonesia"
}

func main(){
	address1 := Address{"Makasssar", "Sulsel", "Indonesia"}
	address2 := &address1 //mereferens ke address1(alamat memorinya)

	address2.City = "Toraja"

	*address2 = Address{"Rantepao", "Sulsel", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	/*
		Kita juga bisa membuat pointer menggunakan fungsi new selain menggunakan operator &
		Fungis new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/
	var address3 *Address= new(Address)
	address3.City = "Makale"
	fmt.Println(address3)

	var alamat = Address{
		City : "Makasssar",
		Province : "Sulsel",
		Country : "",
	}
	// ChangeAddressToIndonesia(alamat)//data asli
	// fmt.Println(alamat)//tidak berubah

	// var alamatPointer *Address = &alamat //cara 1 menggunakan value
	//ChangeAddressToIndonesia(alamat)

	ChangeAddressToIndonesia(&alamat) //cara 2 memakai pointer di func
	fmt.Println(alamat)// berubah


}