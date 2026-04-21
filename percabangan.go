package main 

import "fmt"

func main() {
	nilai := 85

	if nilai >= 90 {
		fmt.Println("Dapat Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Dapat Nilai B")
	} else if nilai >= 65 {
		fmt.Println("Dapat Nilai C")
	} else if nilai >= 50 {
		fmt.Println("Dapat Nilai D")
	} else {
		fmt.Println("Dapat Nilai E")
	}

	//swich case 
	fmt.Println("Contoh switch case")
	hari := "Rabu"
	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari tidak valid")
	}

	//switch case dengan kondisi
	fmt.Println("Contoh switch case dengan kondisi")
	umur := 25
	switch {
	case umur < 13:
		fmt.Println("Anak-anak")
	case umur >= 13 && umur < 20:
		fmt.Println("Remaja")
	case umur >= 20 && umur < 60:
		fmt.Println("Dewasa")
	default:
		fmt.Println("Lansia")
	}
}