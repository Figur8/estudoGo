package main

import "fmt"

//Faça um programa que receba 2 valores e retorne o maior entre eles.

func main() {
	var valueA, valueB float32

	fmt.Println("Write 2 values for compare whats the biggest")
	fmt.Scan(&valueA, &valueB)
	if valueA > valueB {
		fmt.Println("The biggest is ", valueA)
	} else {
		fmt.Println("The biggest is ", valueB)
	}
}
