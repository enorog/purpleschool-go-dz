package main

import (
	"fmt"
)

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

func getConverterInput() (sourceAmount float64, sourceCurrency string, targetCurrency string, err error) {
	for {
		fmt.Print("Введите исходную валюту")
		sourceCurrency, err = getCurrency("")
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
		targetCurrency, err = getCurrency(sourceCurrency)
		if err == nil {
			break
		}
		fmt.Printf("%v: повторите ввод\n", err)
	}
	return
}

func getCurrency(excludedCurrency string) (input string, err error) {
	fmt.Print("(")
	first := true
	for i := 0; i < 3; i++ {
		var currency string
		currency, err = currencyByNumber(i)
		if err != nil {
			panic(err)
		}
		if currency != excludedCurrency {
			if first {
				first = false
			} else {
				fmt.Print(", ")
			}

			fmt.Print(currency)
		}
	}
	fmt.Print("): ")
	_, err = fmt.Scan(&input)
	if err != nil {
		return
	}
	for i := 0; i < 3; i++ {
		var currency string
		currency, err = currencyByNumber(i)
		if err != nil {
			panic(err)
		}
		if currency != excludedCurrency && currency == input {
			return
		}
	}
	err = fmt.Errorf("ошибочный код валюты %s", input)
	return
}

func currencyByNumber(num int) (string, error) {
	switch num {
	case 0:
		return USD, nil
	case 1:
		return EUR, nil
	case 2:
		return RUB, nil
	default:
		return "", fmt.Errorf("ошибочный номер вылюты %d", num)
	}
}

func convertCurrency(sourceAmount float64, sourceCurrency string, targetCurrency string) (targetAmount float64, err error) {
	const USD_TO_EUR = 0.8
	const USD_TO_RUB = 80.2
	const EUR_TO_RUB = 1 / USD_TO_EUR * USD_TO_RUB
	var rate float64
	switch sourceCurrency {
	case USD:
		switch targetCurrency {
		case EUR:
			rate = USD_TO_EUR
		case RUB:
			rate = USD_TO_RUB
		default:
			err = conversionError(sourceCurrency, targetCurrency)
		}
	case EUR:
		switch targetCurrency {
		case USD:
			rate = 1 / USD_TO_EUR
		case RUB:
			rate = EUR_TO_RUB
		default:
			err = conversionError(sourceCurrency, targetCurrency)
		}
	case RUB:
		switch targetCurrency {
		case USD:
			rate = 1 / USD_TO_RUB
		case EUR:
			rate = 1 / EUR_TO_RUB
		default:
			err = conversionError(sourceCurrency, targetCurrency)
		}
	default:
		err = fmt.Errorf("ошибочная исходная валюта %s", sourceCurrency)
	}
	if err != nil {
		return
	}
	targetAmount = sourceAmount * rate
	return
}

func conversionError(sourceCurrency, targetCurrency string) error {
	return fmt.Errorf("ошибочная целевая валюта %s для исходной валюты %s", targetCurrency, sourceCurrency)
}

func main() {
	fmt.Println("============ Калькулятор валют ================")
	for {
		sourceamount, sourceCurrency, targetCurrency, err := getConverterInput()
		if err == nil {
			convertedAmount, err := convertCurrency(sourceamount, sourceCurrency, targetCurrency)
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
