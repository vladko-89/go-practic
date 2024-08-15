package main

import (
	"fmt"
	"strings"
)


func main() {
	exchangeCurs := map[string]map[string]float64{
		"usd": {
			"eur": 0.92,
			"rub": 86.11,
		},
		"rub": {
			"usd": 0.01,
			"eur": 0.01,
		},
		"eur": {
			"usd": 1.05,
			"rub": 0.01,
		},
	}

	var value int
	var targetCurency, currentCurency string

	for {
		currentCurency = getUserInputCurency("Выберите текущую валюту (usd, eur, rub): ")
		if typeOfObject(currentCurency) == "string" && (currentCurency == "usd" || currentCurency == "eur" || currentCurency == "rub") {
			break
		}
		fmt.Println("Введено не подходящее значение")
	}

	for {
		value = getUserInputCount()
		if typeOfObject(value) == "int" && value > 0 {
			break
		}
		fmt.Println("Введено не подходящее значение")
	}

	availibleCurency := ""

	switch currentCurency {
	case "eur":
		availibleCurency = "usd, rub"
	case "usd":
		availibleCurency = "eur, rub"
	case "rub":
		availibleCurency = "usd, eur"
	}

	for {
		targetCurency = getUserInputCurency("Выберите валюту конвертации (" + availibleCurency + "): ")
		if typeOfObject(targetCurency) == "string" && targetCurency != currentCurency && (currentCurency == "usd" || currentCurency == "eur" || currentCurency == "rub") {
			break
		}
		fmt.Println("Введено не подходящее значение")
	}

	fmt.Printf("Вы получите: %.2f", calculateConvertation(&value, &targetCurency, &currentCurency, exchangeCurs))

}

func getUserInputCount() int {
	var val int
	fmt.Print("Введите сумму к конвертации: ")
	fmt.Scan(&val)
	return val
}

func getUserInputCurency(label string) string {
	var val string
	fmt.Print(label)
	fmt.Scan(&val)
	return strings.ToLower(val)
}

func calculateConvertation(
	value *int,
	targetCurency *string,
	currentCurency *string,
	curs map[string]map[string]float64,
) float64 {

	return float64(*value) * curs[*currentCurency][*targetCurency]
}

func typeOfObject(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
