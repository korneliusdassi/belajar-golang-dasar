package main

import "fmt"

var member = map[string]string{
	"nama":           "Kornelius",
	"umur":           "24",
	"jenisKelamin":   "Laki-Laki",
	"agama":          "Kristen",
	"titel":          "S.Kom",
	"pekerjaan":      "Programmer",
	"jabatan":        "Backend Developer",
	"tangggal_masuk": "20 juli 2021",
}

func checkMember(key string) bool {
	_, cetak := member[key] //mengambil key
	return cetak
}

func main() {
	if checkMember("nama") {
		fmt.Println("Member anda terdaftar!")
	} else {
		fmt.Println("Member anda tidak terdaftar!")
	}
}
