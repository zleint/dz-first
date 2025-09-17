package main

import (
	"errors"
	"fmt"
)

const price_USD = 1.0
const price_EUR = 0.93
const price_RUB = 0.011

func main() {
	for {
		var x string
		first_val, second_val, sum := userInput()
		first_val, second_val, sum, err := reviewMaker(first_val, second_val, sum)
		if err != nil {
			fmt.Println("Error:\n", err)
			continue
		}
		result := converter(first_val, second_val, sum)
		fmt.Printf("Результат конвертации %v из %v равен %v\n", first_val, second_val, result)
		fmt.Println("Хотите попробовать снова? y/n")
		fmt.Scan(&x)
		if x == "y" {
			continue
		}
		break //end
	} //break
}
func userInput() (string, string, float64) {
	var first_val string
	var second_val string
	var sum float64
	fmt.Println("Введите первую валюту (USD, RUB, EUR)")
	fmt.Scan(&first_val)
	switch first_val {
	case "USD":
		fmt.Printf("Введите вторую валюту(RUB, EUR)\n")
	case "RUB":
		fmt.Printf("Введите вторую валюту(USD, EUR)\n")
	case "EUR":
		fmt.Printf("Введите вторую валюту(RUB, USD)\n")
	}
	fmt.Scan(&second_val)
	fmt.Printf("Введите сколько %s вы хотите конвертировать\n", first_val)
	fmt.Scan(&sum)
	return first_val, second_val, sum
}
func reviewMaker(first_val string, second_val string, sum float64) (string, string, float64, error) {
	switch {
	case first_val != "USD" && first_val != "RUB" && first_val != "EUR":
		return first_val, second_val, sum, errors.New("неправильно введена первая валюта")
	case second_val != "USD" && second_val != "RUB" && second_val != "EUR":
		return first_val, second_val, sum, errors.New("неправильно введена вторая валюта")
	case sum <= 0:
		return first_val, second_val, sum, errors.New("неправильно введена сумма")
	default:
		return first_val, second_val, sum, nil
	}
}
func converter(first_val string, second_val string, sum float64) float64 {
	var x, y float64
	switch first_val {
	case "USD":
		x = price_USD
	case "RUB":
		x = price_RUB
	case "EUR":
		x = price_EUR
	}

	switch second_val {
	case "USD":
		y = price_USD
	case "RUB":
		y = price_RUB
	case "EUR":
		y = price_EUR
	}
	result := (sum * y) / x
	return result
}
