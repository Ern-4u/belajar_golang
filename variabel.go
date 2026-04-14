package main
import "fmt"

func main() {
	// jika menggunakan "=" maka variabel harus di deklarasikan tipe datanya terlebih dahulu
	var name string = "day 1 isa belajar golang"
	fmt.Println(name)
	// jika menggunakan	 ":=" maka variabel tidak usah di deklarasikan tipe databya,
	variabel := "day 1 isa belajar golang"
	fmt.Println(variabel)
}