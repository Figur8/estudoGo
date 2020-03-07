package main

import "fmt"

//Faça um programa que verifique se um número é impar
func main() {
	var valueA int
	fmt.Println("Write a number: ")
	fmt.Scan(&valueA)
	if (valueA % 2) == 0 {
		fmt.Println("it's even")
	} else {
		fmt.Println("it's odd")
	}
}
