package main

import (
	"fmt"
)

const price_USD = 1.0
const price_EUR = 0.93
const price_RUB = 0.011

func main() {
	for {
		first_val := inputFirstVal()
		second_val := inputSecondVal(first_val)
		sum := inputSum(first_val)
		var x string
		result := converter(first_val, second_val, sum)
		fmt.Printf("Результат конвертации %v из %v равен %v\n", first_val, second_val, result)
		fmt.Println("Хотите попробовать снова? y/n")
		fmt.Scan(&x)
		if x == "y" {
			continue
		}
		break
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
func inputFirstVal() string {
	for {
		var first_val string
		fmt.Println("Введите первую валюту (USD, RUB, EUR)")
		fmt.Scan(&first_val)
		if first_val != "USD" && first_val != "RUB" && first_val != "EUR" {
			fmt.Println("Ошибка, попробуйте снова")
			continue
		}
		return first_val
	}
}
func inputSecondVal(first_val string) string {
	for {
		var second_val string
		switch first_val {
		case "USD":
			fmt.Printf("Введите вторую валюту(RUB, EUR)\n")
		case "RUB":
			fmt.Printf("Введите вторую валюту(USD, EUR)\n")
		case "EUR":
			fmt.Printf("Введите вторую валюту(RUB, USD)\n")
		}
		fmt.Scan(&second_val)
		if second_val != "USD" && second_val != "RUB" && second_val != "EUR" {
			fmt.Println("Ошибка, попробуйте снова")
			continue
		}
		return second_val
	}
}
func inputSum(first_val string) float64 {
	for {
		var sum float64
		fmt.Printf("Введите сколько %s вы хотите конвертировать\n", first_val)
		fmt.Scan(&sum)
		if sum <= 0 {
			fmt.Println("Ошибка, попробуйте снова")
			continue
		}
		return sum
	}
}
