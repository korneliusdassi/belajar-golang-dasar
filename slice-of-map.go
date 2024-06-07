package main

import "fmt"

func main() {
	//slice berisi map
	students := []map[string]string{
		{"name": "lius", "skill": "jago backend"},
		{"name": "dassi", "skill": "jago frontend"},
		{"name": "kornel", "skill": "fullstack"},
	}
	// fmt.Println(students)

	for _, student := range students {
		fmt.Println(student["name"], "-", student["skill"])
	}
}
