package main

import "fmt"

/*
	contoh kasus pointer menggunakan method
*/

type Gamers struct {
	Name  string
	Games []string
}

func (gamer *Gamers) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamers{Name: "Lius"}
	gamer.AddGame("pubg")
	gamer.AddGame("free fire")
	gamer.AddGame("mobile legend")
	gamer.AddGame("coc")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
