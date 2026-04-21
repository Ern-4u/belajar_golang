package main

import "fmt"

//fungsi dasar tanpa parameter dan tanpa return value
func dasar() {
	fmt.Println("Halo, selamat pagi!")
}

//fungsi dengan parameter dan tanpa return value
func nama(name string) {
	fmt.Printf("Halo, %s! Selamat datang di dunia pemrograman Go!\n", name)
}

//fungsi dengan parameter dan return value
func tambah(a, b int) int {
	return a + b
}

func main() {
	//memanggil fungsi dasar
	dasar()
	//memanggil fungsi nama dengan parameter
	nama("Isa")
	//memanggil fungsi tambah dengan parameter dan menyimpan hasilnya
	result := tambah(5, 3)
	fmt.Printf("Hasil penjumlahan 5 + 3 = %d\n", result)
}