package main

import "fmt"

func main () {
	// for tradisional
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan for tradisional ke : ", i)
	}

	//for pengganti while
	fmt.Println("Perulangan for pengganti while")
	bensin := 5
	for bensin > 0 {
		fmt.Println("Bensin tersisa : ", bensin)
		bensin--
	}
	fmt.Println("Bensin habis")

	//for pengganti do while
	fmt.Println("Perulangan for pengganti do while")
	angka := 0
	for {
		fmt.Println("Angka : ", angka)
		angka++
		if angka > 5 {
			break
		}	
	}

	//for range
	fmt.Println("Perulangan for range")
	keranjang := []string{"apel", "jeruk", "mangga"}
	for index, buah := range keranjang {
		fmt.Println("Index :", index," Buah :", buah)
	}


	//di cara yang sebelumnya, kamu bertanya bagaimana juka menghapus slice tanpa membuat
	//index nya berantakan.
	//cara yang paling disarankan di golang bukanlah memotong slice asli melainkan membuat sclice baru yang kosong
	//lalu menggunakan fo range dan if untuk memfilter data
	//contoh kita ingin menghapus buah "jeruk" dari keranjang
	fmt.Println("Menghapus buah jeruk dari keranjang")
	keranjangLama := []string{"apel", "jeruk", "mangga"}
	keranjangBaru := []string{}
	for _, buah := range keranjangLama {
		if buah != "jeruk" {
			keranjangBaru = append(keranjangBaru, buah)
		}
	}
	fmt.Println("Keranjang baru : ", keranjangBaru)

	



}