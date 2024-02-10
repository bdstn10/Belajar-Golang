package main

import "fmt"

func main() {
	// while loop syntax in Golang
	var i int = 0
	for {
		fmt.Println("Berhitung ", i)

		i++
		if i > 5 {
			break
		}

	}
}
