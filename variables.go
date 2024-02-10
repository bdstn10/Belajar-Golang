package main

import "fmt"

func main() {
	// Membuat array daftar_siswa dengan ukuran yang tidak dibatasi
	var daftar_siswa = [...]string{
		"Budi Setiawan",
		"Adifa Kyla",
		"Ramandhika",
	}

	fmt.Println(daftar_siswa[0], daftar_siswa[1], daftar_siswa[2])
}
