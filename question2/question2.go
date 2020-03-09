package main

import "fmt"

func main() {
	//Faça um programa que receba o valor do quilo de um produto e
	//a quantidade de quilos do produto consumida calculando o valor final a ser pago.
	var weightCost float32
	var consumedWeight float32
	var payValue float32
	fmt.Println("How much cost the product weight?")
	fmt.Scan(&weightCost)
	fmt.Println("what's your consumeWeight?")
	fmt.Scan(&consumedWeight)
	payValue = weightCost * consumedWeight
	fmt.Println(payValue)
}
