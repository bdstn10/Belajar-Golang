package main

import "fmt"

func main() {
	var warnaLampu string = "hijau"

	switch warnaLampu {
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
