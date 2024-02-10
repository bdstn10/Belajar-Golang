package main

import "fmt"

func main() {
	var warnaLampu string = "biru"

	switch warnaLampu {
	// Switch case multiple case
	case "jingga", "ungu", "biru":
		fmt.Println("Mungkin itu warna pelangi")
	case "merah":
		{
			fmt.Println("Mohon berhenti!")
		}
	case "kuning":
		fmt.Println("Hati-hati")
		fmt.Println("Jangan nekat")
	case "hijau":
		fmt.Println("Silahkan jalan")
	default:
		{
			fmt.Println("Tidak tahu!")
		}
	}
}
