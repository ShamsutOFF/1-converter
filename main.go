package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	usdCourse := map[string]float64{"EUR": 0.85, "RUB": 85.00, "USD": 1}
	fmt.Println("___ Добро пожаловать в конвертор валют! ___")
	var cur1, cur2 string
	var sum int
	for {
		cur, err := readFirstCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		cur1 = cur
		break
	}
	fmt.Println("Первая валюта: ", cur1)

	for {
		s, err := readSum()
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum = s
		break
	}
	fmt.Println("Сумма для конвертации: ", sum)

	for {
		cur, err := readSecondCurrency(cur1)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cur2 = cur
		break
	}
	fmt.Println("Вторая валюта: ", cur2)
	// Вызов функции конвертации
	result := convert(float64(sum), cur1, cur2, usdCourse)
	fmt.Printf("Результат конвертации: %.2f %s = %.2f %s\n", float64(sum), cur1, result, cur2)
}

func convert(sum float64, fromV, toV string, courses map[string]float64) float64 {
	usdAmount := sum / courses[fromV]
	return usdAmount * courses[toV]
}

func readFirstCurrency() (string, error) {
	fmt.Print("Введите исходную валюту(USD, EUR или RUB): ")
	input, err := readInput()
	if err != nil {
		return "", err
	}

	input = strings.ToUpper(strings.TrimSpace(input))
	if input != "USD" && input != "EUR" && input != "RUB" {
		return "", errors.New("вы ввели не верную валюту! Попробуйте снова")
	}
	return input, nil
}

func readSecondCurrency(firstCur string) (string, error) {
	fmt.Print("Введите целевую валюту(USD, EUR или RUB): ")
	input, err := readInput()
	if err != nil {
		return "", err
	}

	input = strings.ToUpper(strings.TrimSpace(input))
	if input != "USD" && input != "EUR" && input != "RUB" {
		return "", errors.New("вы ввели не верную валюту! Попробуйте снова")
	}
	if input == firstCur {
		return "", errors.New("целевая валюта не должна совпадать с исходной! Попробуйте снова")
	}
	return input, nil
}

func readSum() (int, error) {
	fmt.Print("Введите сумму для конвертации: ")
	input, err := readInput()
	if err != nil {
		return 0, err
	}

	var sum int
	_, err = fmt.Sscanf(input, "%d", &sum)
	if err != nil {
		return 0, errors.New("вы ввели не верную сумму! Попробуйте снова")
	}
	if sum <= 0 {
		return 0, errors.New("сумма должна быть положительным числом! Попробуйте снова")
	}
	return sum, nil
}

// Новая функция для чтения ввода с очисткой буфера
func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("ошибка чтения ввода")
	}
	return strings.TrimSpace(input), nil
}
