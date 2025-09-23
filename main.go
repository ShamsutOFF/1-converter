package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var operations = map[string]func([]int) float64{
	"AVG": func(nums []int) float64 {
		var result float64
		for _, num := range nums {
			result += float64(num)
		}
		result /= float64(len(nums))
		return result
	},
	"SUM": func(nums []int) float64 {
		var result float64
		for _, num := range nums {
			result += float64(num)
		}
		return result
	},
	"MED": func(nums []int) float64 {
		var result float64
		sort.Ints(nums)
		if len(nums)%2 == 0 {
			i := len(nums) / 2
			result = (float64(nums[i]) + float64(nums[i-1])) / 2.0
		} else {
			result = float64(nums[len(nums)/2])
		}
		return result
	},
}

func main() {
	fmt.Println("___ Добро пожаловать в калькулятор ___")
	var operation string
	for {
		op, err := readOperation()
		if err != nil {
			fmt.Println(err)
			continue
		}
		operation = op
		break
	}

	fmt.Println("Операция:", operation)
	var nums []int
	for {
		ints, err := readNumbers()
		if err != nil {
			fmt.Println(err)
			continue
		}
		nums = ints
		break
	}
	fmt.Println("Вы ввели: ", nums)
	result := operations[operation](nums)

	fmt.Println("Результат операции: ", result)
}

func readOperation() (string, error) {
	fmt.Print("Введите целевую операцию (AVG - среднее, SUM - сумму или MED - медиану): ")
	input, err := readInput()
	if err != nil {
		return "", err
	}

	input = strings.ToUpper(strings.TrimSpace(input))
	if input != "AVG" && input != "SUM" && input != "MED" {
		return "", errors.New("вы ввели не верную операцию! Попробуйте снова")
	}

	return input, nil
}

func readNumbers() ([]int, error) {
	fmt.Print("Введите числа через запятую: ")
	input, err := readInput()
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	slStr := strings.Split(input, ",")
	var slInt = make([]int, len(slStr))
	for i, v := range slStr {
		v = strings.TrimSpace(v)
		atoi, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		slInt[i] = atoi
	}
	return slInt, nil
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("ошибка чтения ввода")
	}
	return strings.TrimSpace(input), nil
}
