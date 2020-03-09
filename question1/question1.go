package main

import (
	"fmt"
	"os"
)

func sellBusinessRule(change int) {
	if change >= 0 {
		fmt.Println("Your change is: ", change)
	} else {
		fmt.Println("Insuficient money")
	}
}
func sellingProcess() {
	var paidValue, productPrice int
	fmt.Println("Faça um programa que receba um valor que é o valor pago, um segundo valor que é o preço do produto e retorne o troco a ser dado.")
	fmt.Println("Enter the paid value and product price respectively: ")
	fmt.Scan(&paidValue, &productPrice)
	change := paidValue - productPrice
	sellBusinessRule(change)
	os.Exit(0)
}
func main() {
	sellingProcess()
}
