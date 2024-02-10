package main

import "fmt"

func main() {
	if warnaLampu := "hijau"; warnaLampu == "merah" {
		fmt.Println("Mohon berhenti")
	} else if warnaLampu == "kuning" {
		fmt.Println("Mohon Hati-Hati")
	} else {
		fmt.Println("Silahkan Jalan")
	}
}
