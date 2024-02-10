package main

import "fmt"

func main() {
	// Membuat array daftar_siswa dengan ukuran 3 elemen
	var daftar_siswa [3]string
	daftar_siswa[0] = "Budi Setiawan"
	daftar_siswa[1] = "Adifa Kyla"
	daftar_siswa[2] = "Ramandhika"

	fmt.Println(daftar_siswa[0], daftar_siswa[1], daftar_siswa[2])
}
