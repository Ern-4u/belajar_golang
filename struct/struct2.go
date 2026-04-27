package main

import "fmt"

type Kucing struct {
	Nama  string
	Warna string
}

// Ini adalah Method. Perhatikan tambahan (k Kucing) sebelum nama fungsi.
// Artinya: Fungsi 'Bersuara' ini HANYA milik struct Kucing.
func (k Kucing) Bersuara() {
	fmt.Println(k.Nama, "berkata: Meowww!")
}

func main() {
	kucingKu := Kucing{Nama: "Oyen", Warna: "Oranye"}

	fmt.Println("Nama Kucing:", kucingKu.Nama)
	fmt.Println("Warna Kucing:", kucingKu.Warna)
	
	// Memanggil method menggunakan tanda titik
	kucingKu.Bersuara() // Output: Oyen berkata: Meowww!
}