package main

import "fmt"

func main() {
	var angka1 int = 10
	angka2 := 20
	//operator pertambahan menggunakan simbol +
	haslit_tambah := angka1 + angka2
	fmt.Println("hasil pertambahan : ", haslit_tambah)

	//operator pengurangan menggunakan simbol - 
	hasil_pengurangan := angka1 - angka2
	fmt.Println("hasil pengurangan : ", hasil_pengurangan)

	//operator perkalian menggunakan simbol *
	hasil_perkalian := angka1 * angka2
	fmt.Println("hasil perkalian : ", hasil_perkalian)

	//operator pembagian menggunakan simbol /
	hasil_pembagian := angka1 / angka2
	fmt.Println("hasil pembagian : ", hasil_pembagian)

	//operator modulus menggunakan simbol %
	hasil_modulus := angka1 % angka2
	fmt.Println("hasil modulus : ", hasil_modulus)

}