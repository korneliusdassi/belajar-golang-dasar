package main

import "fmt"

/*
Embedded struct
*/
type Member struct {
	ID       int
	Nama     string
	IsActive bool
}

type Group struct {
	Name        string
	Admin       Member   //admin tipenya adalah Member
	Members     []Member //artinya group boleh punya banyak anggota jadi anggotanya bertipe slice of Member
	IsAvailable bool
}

//struct sebagai parameter function
func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Members))

	fmt.Println(" Member name:")
	for _, user := range group.Members {
		fmt.Println(user.Nama)
	}
}

func main() {
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

	users := []Member{user1, user2}
	group := Group{"Gamer", user1, users, true}
	displayGroup(group)
}
