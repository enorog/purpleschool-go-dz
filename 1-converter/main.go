package main

import "fmt"

func getConverterInput() (sourceAmount float64, sourceCurrency string, targetCurrency string) {
	fmt.Print("Введите исходную валюту: ")
	fmt.Scanf("%3s\n", &sourceCurrency)
	fmt.Print("Введите сумму ковертации: ")
	fmt.Scanf("%f\n", &sourceAmount)
	fmt.Printf("Веедите целевую валюту: ")
	fmt.Scanf("%3s\n", &targetCurrency)
	return
}

func convertCurrency(sourceamount float64, sourceCurrency string, targetCurrency string) float64 {
	return 0
}

func main() {
	const USD_TO_EUR = 0.8
	const USD_TO_RUB = 80.2
	const EUR_TO_RUB = 1 / USD_TO_EUR * USD_TO_RUB

	fmt.Println("Constants:")
	fmt.Printf("USD_TO_EUR: %f\n", USD_TO_EUR)
	fmt.Printf("USD_TO_RUB: %f\n", USD_TO_RUB)
	fmt.Printf("EUR_TO_RUB: %f\n", EUR_TO_RUB)

	sourceamount, sourceCurrency, targetCurrency := getConverterInput()

	convertedAmount := convertCurrency(sourceamount, sourceCurrency, targetCurrency)
	fmt.Printf("Сумма в целевой валюте: %.2f\n", convertedAmount)
}
