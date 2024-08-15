package main

import (
	"fmt"
	"strconv"
	"strings"
)

const AVG = "avg"
const SUM = "sum"
const MED = "med"

func main() {
	var operation string
	var inputNums string

	for {
		fmt.Print("Введите операцию (avg, sum, med): ")
		fmt.Scan(&operation)
		if typeOfObject(operation) == "string" && strings.ToLower(operation) == AVG || strings.ToLower(operation) == SUM || strings.ToLower(operation) == MED {
			break
		}
		fmt.Println("Можно использовать только одну из операций: AVG, SUM, MED")

	}

	fmt.Print("Введите числа через запятую (1,4,5): ")
	fmt.Scan(&inputNums)

	nums := make([]int, len(strings.Split(inputNums, ",")))

	for i, val := range strings.Split(inputNums, ",") {
		temp, err := strconv.Atoi(val)

		if err != nil {
			nums[i] = temp

		} else {
			continue

		}

	}

	fmt.Printf("Ответ: %d", calculate(operation, nums))

}

func calculate(oper string, list []int) int {
	var res int

	switch oper {
	case SUM:
		for _, item := range list {
			res += item
		}
	case MED:
		if len(list)%2 == 0 {
			res = int((list[len(list)/2-1] + list[len(list)/2+1]) / 2)
		} else {
			res = int(list[len(list)/2])
		}
	case AVG:
		temp := 0
		for _, item := range list {

			temp += item

		}
		res = int(temp / len(list))

	}
	return res
}

func typeOfObject(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
