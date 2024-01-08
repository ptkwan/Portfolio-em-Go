package main

import "fmt"

//Caso múltiplo de 3 você fala PIN
//Se for múltiplo de 5, PAN
//Se for os dois, PIN PAN

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PIN PAN", i)
		} else if i%3 == 0 {
			fmt.Println("PIN", i)
		} else if i%5 == 0 {
			fmt.Println("PAN", i)
		}
	}
}
