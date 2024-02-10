package main

import "fmt"

func main() {
	var warnaLampu string = "kuning"

	// if else like switch case usage
	switch {
	case len(warnaLampu) > 5:
		fmt.Println("Ada cukup banyak karakter di variabel tersebut")
	case warnaLampu == "merah":
		{
			fmt.Println("Mohon berhenti!")
		}
	case warnaLampu == "kuning":
		fmt.Println("Hati-hati")
		fmt.Println("Jangan nekat")
	case warnaLampu == "hijau":
		fmt.Println("Silahkan jalan")
	default:
		{
			fmt.Println("Switch case selesai")
		}
	}
}
