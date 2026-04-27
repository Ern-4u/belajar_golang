package main

import (
	"fmt"
	"net/http"
)

// 1. Handler untuk Halaman Utama
func halamanUtama(w http.ResponseWriter, r *http.Request) {
	// Memberi tahu browser bahwa kita mengirim file HTML, bukan teks biasa
	w.Header().Set("Content-Type", "text/html")
	
	// Menulis tag HTML langsung ke dalam ResponseWriter
	fmt.Fprintf(w, `
		<html>
			<head><title>Web Golang</title></head>
			<body>
				<h1> Selamat Datang di Web Golang!</h1>
				<p>Ini adalah halaman utama.</p>
				<a href="/profil">Lihat Profil Saya</a>
			</body>
		</html>
	`)
}

// 2. Handler untuk Halaman Profil
func halamanProfil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	fmt.Fprintf(w, `
		<html>
			<head><title>Profil - Web Golang</title></head>
			<body>
				<h1>Halaman Profil</h1>
				<p>Halo, saya adalah calon programmer Golang handal.</p>
				<a href="/">Kembali ke Beranda</a>
			</body>
		</html>
	`)
}

func main() {
	// 3. Mendaftarkan Rute (Routing)
	http.HandleFunc("/", halamanUtama)         // Jika URL adalah localhost:8080/
	http.HandleFunc("/profil", halamanProfil)  // Jika URL adalah localhost:8080/profil

	// 4. Menyalakan Server
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}