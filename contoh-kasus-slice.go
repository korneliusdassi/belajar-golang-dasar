package main

import "fmt"

func main() {
	/*
		1. Hitung nilai rata2 dari array berikut
		penye: total dari jumlah array / panjang array
		2. Cari nilai >=  90 mengunakan slice
	*/
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total int
	for _, score := range scores {
		total += score
	}

	length := len(scores)
	average := float64(total) / float64(length)

	fmt.Println("Nilai rata-rata dari array diatas adalah: ", average)

	fmt.Println("---------------------------------------------------------")
	var goodScores []int
	for _, skor := range scores {
		if skor >= 90 {
			goodScores = append(goodScores, skor)
		}
	}
	fmt.Println("nilai >=  90 adalah: ", goodScores)
}
