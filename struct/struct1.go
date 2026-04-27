package main

import "fmt"

// 1. MEMBUAT CETAKAN (Blueprint)
// Biasanya ditulis di luar func main
type Pelanggan struct {
	Nama   string
	Umur   int
	IsVIP  bool
}

func main() {
	// 2. MEMBUAT OBJEK DARI CETAKAN
	// Cara 1: Menulis nama field-nya (Paling disarankan agar tidak tertukar)
	pelanggan1 := Pelanggan{
		Nama:  "Budi Santoso",
		Umur:  25,
		IsVIP: false,
	}

	// 3. MENGAKSES DATA DI DALAM STRUCT
	// Gunakan tanda titik (.)
	fmt.Println("Nama Pelanggan:", pelanggan1.Nama)
	
	if pelanggan1.IsVIP {
		fmt.Println("Berikan diskon khusus untuk Budi!")
	}
}