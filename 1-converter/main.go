package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.8
	const USD_TO_RUB = 80.2
	const EUR_TO_RUB = 1 / USD_TO_EUR * USD_TO_RUB

	fmt.Println("Constants:")
	fmt.Printf("USD_TO_EUR: %f\n", USD_TO_EUR)
	fmt.Printf("USD_TO_RUB: %f\n", USD_TO_RUB)
	fmt.Printf("EUR_TO_RUB: %f\n", EUR_TO_RUB)
}
