package main

import "fmt"

func main(){

	points := 10
	if points > 7 {
		switch points{
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	}else{
		if points == 5{
			fmt.Println("not bad")
		}else if points == 3{
			fmt.Println("keep trying")
		}else{
			fmt.Println("you can do it")
			if points ==0 {
				fmt.Println("try harder")
			}
		}
	}

	// var point = 6
	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// 	fallthrough
	// case point < 5:
	// 	fmt.Println("you need to learn more")
	// default:
	// {
	// 	fmt.Println("not bad")
	// 	fmt.Println("you need to learn more")
	// }
	// }

	// color := ""

	// switch color{
	// case "merah":
	// 	fmt.Println("lanjut")
	// case "biru":
	// 	fmt.Println("lanjut")
	// default:
	// fmt.Println("warna tdk ada")

	// }
}