package main

import "fmt"

const usdEur = 0.92
const usdRub = 86.11
const eurRub = usdRub / usdEur

func main() {
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

	fmt.Printf("Вы получите: %.2f", calculateConvertation(value, targetCurency, currentCurency))

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
	return val
}

func calculateConvertation(value int, targetCurency, currentCurency string) float64 {
	if targetCurency == "usd" && currentCurency == "rub" {
		return float64(value) / usdRub
	} else if targetCurency == "eur" && currentCurency == "rub" {
		return float64(value) / eurRub
	} else if targetCurency == "rub" && currentCurency == "usd" {
		return float64(value) * eurRub
	} else if targetCurency == "rub" && currentCurency == "eur" {
		return float64(value) * eurRub
	} else if targetCurency == "usd" && currentCurency == "eur" {
		return float64(value) / usdEur
	}
	return float64(value) * usdEur
}

func typeOfObject(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
