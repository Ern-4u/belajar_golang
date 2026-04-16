package main

import "fmt"

func main() {
	//output "Hello, World!" ke layar menggunakan fungsi fmt.Println()
	fmt.Println("Hello, World!")

	//input "nama" dari user menggunakan fungsi fmt.Scanln()
	var nama string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&nama)

	//output "Hello, nama!" ke layar menggunakan fungsi fmt.Println()
	fmt.Println("Hello, ", nama)

}