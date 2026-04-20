package main

import "fmt"

func main() {
	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println("slice1 : ", slice1)

	slice2 := []string{"golang", "python", "java"}
	fmt.Println("slice2 : ", slice2)

	//mwngakses elemen slice menggunakan index
	fmt.Println("elemen pertama slice1 : ", slice1[0])
	fmt.Println("elemen kedua slice2 : ", slice2[1])

	//menambahkan data ke dalam slice menggunakan fungsi append()
	slice1 = append(slice1, 6)
	fmt.Println("slice1 setelah ditambahkan : ", slice1)
	slice2 = append(slice2, "javascript")
	fmt.Println("slice2 setelah ditambahkan : ", slice2)

	//mengedit elemen slice mrnggunakan fungsi index
	slice1[0] = 10
	fmt.Println("slice1 setelah diedit : ", slice1)
	slice2[1] = "ruby"
	fmt.Println("slice2 setelah diedit : ", slice2)

	//menghapus elemen slice menggunakan fungsi copy() dan slicing
	fmt.Println("isi slice yang belum dihapus :", slice1)
	fmt.Println("isi slice yang belum dihapus :", slice2)
	slice1 = append(slice1[:2], slice1[3:]...)
	fmt.Println("slice1 setelah dihapus : ", slice1)
	slice2 = append(slice2[:1], slice2[2:]...)
	fmt.Println("slice2 setelah dihapus : ", slice2)
	
}