package main

import (
	"fmt"
)

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

type conversionRateMap = map[string]float64

func getConverterInput(rates conversionRateMap) (sourceAmount float64, sourceCurrency string, targetCurrency string, err error) {

	for {
		fmt.Print("Введите исходную валюту")
		sourceCurrency, err = getCurrency(rates, "")
		if err == nil {
			break
		}
		fmt.Printf("%v: повторите ввод\n", err)
	}
	for {
		fmt.Print("Введите сумму конвертации: ")
		_, err = fmt.Scanf("%f\n", &sourceAmount)
		if err == nil && sourceAmount < 0 {
			err = fmt.Errorf("сумма ковертации должа быть больше нуля")
		}

		if err == nil {
			break
		}
		fmt.Printf("%v: повторите ввод\n", err)
	}
	for {
		fmt.Print("Веедите целевую валюту")
		targetCurrency, err = getCurrency(rates, sourceCurrency)
		if err == nil {
			break
		}
		fmt.Printf("%v: повторите ввод\n", err)
	}
	return
}

func getCurrency(currencies conversionRateMap, excludedCurrency string) (input string, err error) {
	fmt.Printf("(%s): ", getCurrencyListString(currencies, excludedCurrency))
	_, err = fmt.Scan(&input)
	if err != nil {
		return
	}
	_, exists := currencies[input]
	if !exists || input == excludedCurrency {
		err = fmt.Errorf("ошибочный код валюты %s", input)
	}
	return
}

func getCurrencyListString(currencies conversionRateMap, excludedCurrency string) (result string) {
	first := true
	for currency := range currencies {
		if currency != excludedCurrency {
			if first {
				first = false
			} else {
				result += ", "
			}

			result += fmt.Sprint(currency)
		}
	}
	return
}

func convertCurrency(sourceAmount float64, sourceCurrency string, targetCurrency string, usdRates conversionRateMap) (targetAmount float64, err error) {
	rate1, ok := usdRates[sourceCurrency]
	if !ok {
		err = conversionError(sourceCurrency, targetCurrency)
		return
	}
	rate2, ok := usdRates[targetCurrency]
	if !ok {
		err = conversionError(sourceCurrency, targetCurrency)
		return
	}
	targetAmount = sourceAmount * (1 / rate1 * rate2)
	return
}

func conversionError(sourceCurrency, targetCurrency string) error {
	return fmt.Errorf("ошибочная целевая валюта %s для исходной валюты %s", targetCurrency, sourceCurrency)
}

func main() {
	fmt.Println("============ Калькулятор валют ================")
	usdRates := conversionRateMap{
		EUR: 0.8,
		RUB: 80.2,
		USD: 1.0,
	}
	for {

		sourceamount, sourceCurrency, targetCurrency, err := getConverterInput(usdRates)
		if err == nil {
			convertedAmount, err := convertCurrency(sourceamount, sourceCurrency, targetCurrency, usdRates)
			if err != nil {
				fmt.Printf("Ошбика ковертации: %v", err)
			} else {
				fmt.Printf("Сумма в целевой валюте: %.2f\n", convertedAmount)
			}
		}
		var cont string
		fmt.Printf("Хотите продолжить (y/n)? ")
		_, err = fmt.Scan(&cont)
		if err != nil {
			fmt.Println(err)
			break
		}
		if cont != "y" {
			break
		}
	}

}
