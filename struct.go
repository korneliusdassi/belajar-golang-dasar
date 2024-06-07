package main

import (
	"fmt"
)

/*
	deklarasi struct
	struct merupakan kumpulan dari field
*/

type Member struct {
	ID       int
	Nama     string
	IsActive bool
}

/*
struct method/function
*/
func (member Member) sayHello(nama string) {
	fmt.Println("Hello", nama, "nama saya", member.Nama)
}

func (a Member) hello() {
	fmt.Println("Hello", a.Nama)
}

/*
	struct sebagai parameter dari function
*/

// func displayMember(member Member) string {
// 	result := fmt.Sprintf("Nama : %s", member.Nama)
// 	return result
// }

func main() {
	//cara instans mengisi struct
	user1 := Member{
		ID:       1,
		Nama:     "Kornelius",
		IsActive: true,
	}
	user2 := Member{
		ID:       2,
		Nama:     "Dassi",
		IsActive: false,
	}

	// fmt.Println(user)
	user1.sayHello("Lius")
	user2.hello()

	// displayMember1 := displayMember(user1)
	// displayMember2 := displayMember(user2)

	// fmt.Println(displayMember1)
	// fmt.Println(displayMember2)
}
