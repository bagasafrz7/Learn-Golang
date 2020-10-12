package main

import (
	"fmt"
	"log"
)

func cetak_gambar(number int) {
	if number % 2 == 0 {
		log.Println("Hallo, The input parameter must be an odd number!")
	} else {
		for i:= 0; i < number; i++ {
			for j:= 0; j < number; j++ {
				if j < 1 || j == number-1 || i == number/2 {
					fmt.Print("* ") 
				} else {
					fmt.Print("= ")
				}
			}
			fmt.Println("\n")
		}
	}
}

func main() {
	cetak_gambar(5)
}