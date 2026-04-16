package main
import "fmt"
import "strings"

func main() {
	var name string = "muhammad isa	irawanto"
	nama_panggilan := "isa"
	fmt.Println(name)
	fmt.Println(nama_panggilan, "\n")

	// jika string menggunaka tanda kutip dua ("") maka akan menghasilkan string biasa
	// tetapi jika string menggunakan tanda kutip backtick (``) maka akan menghasilkan string literal
	fmt.Println("string biasa :")
	// contoh string biasa
	string_biasa := "muhammad isa irawanto saya sedang belajar golang"
	fmt.Println(string_biasa, "\n")


	// contoh string literal
	fmt.Println("string literal :")
	string_literal := `muhammad isa irawanto
saya sedang belajar golang`
	fmt.Println(string_literal, "\n")


	// penggabungan dua string menggunakan operator + 
	fmt.Println("penggabungan string :")
	// contoh
	string1 := "muhammad isa irawanto"
	string2 := "saya sedang belajar golang"
	string3 := string1 + " " + string2
	fmt.Println(string3, "\n")

	// mengubah isi string menjadi huruf kapital menggunakan fungsi strings.ToUpper()
	fmt.Println("mengubah string menjadi huruf kapital :")
	// contoh
	string_kecil := "muhammad isa irawanto"
	string_besar := strings.ToUpper(string_kecil)
	fmt.Println(string_besar, "\n")

	// mengubah isi string menjadi huruf kecil menggunakan fungsi strings.ToLower()]
	fmt.Println("mengubah string menjadi huruf kecil : ")
	// contoh
	string_besar2 := "MUHAMMAD ISA IRAWANTO"
	string_kecil2 := strings.ToLower(string_besar2)
	fmt.Println(string_kecil2, "\n")

	// mengubah isi string agar menjadi huruf kapital di setiap awal kata menggunakan fungsi strings.Title()
	fmt.Println("mengubah string menjadi huruf kapital di setiap awal kata :")
	// contoh
	string_kecil3 := "muhammad isa irawanto"
	string_title := strings.Title(string_kecil3)
	fmt.Println(string_title, "\n")

	// menghapus spasi di awal dan akhir string menggunakan fungsi strings.TrimSpace()
	fmt.Println("menghapus spasi di awal dan akhir string :")
	// contoh
	string_spasi := "   muhammad isa irawanto   "
	string_bersih := strings.TrimSpace(string_spasi)
	fmt.Println(string_bersih, "\n")

	// mengubah isi kata di dalam string menggunakan fungsi strings.Replace()
	fmt.Println("mengubah isi kata di dalam string :")
	// contoh
	string_kecil4 := "muhammad isa irawanto"
	string_baru := strings.Replace(string_kecil4, "isa", "irawan", -1)
	fmt.Println(string_baru, "\n")

	//mengecek berapa banyak furuf di dalam string menggunakan fungsi len()
	fmt.Println("mengecek berapa banyak huruf di dalam string :")
	// contoh
	string_kecil5 := "muhammad isa irawanto"
	panjang_string := len(string_kecil5)
	fmt.Println(panjang_string, "\n")

	// mengecek apakah string mengandung kata tertentu menggunakan fungsi strings.Contains()
	fmt.Println("mengecek apakah string mengandung kata tertentu :")
	// contoh
	string_kecil6 := "muhammad isa irawanto"
	cek_kata := strings.Contains(string_kecil6, "isa")
	fmt.Println(cek_kata, "\n")

	// mengecek berapa jumlah  kata di dalam string menggunakan fungsi strings.Count()
	fmt.Println("mengecek jumlah kata di dalam string :")
	// contoh
	string_kecil7 := "muhammad isa irawanto"
	jumlah_kata := strings.Count(string_kecil7, " ")
	fmt.Println(jumlah_kata + 1, "\n")

	//memecah isi string menjadi beberapa bagian menggunakan fungsi string.Split()
	fmt.Println("memecah isis string menjadi beberapa bagian : 	")
	// contoh
	string_kecil8 := "muhammad isa irawanto"
	string_split := strings.Split(string_kecil8, ",")
	fmt.Println(string_split, "\n")

}