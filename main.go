package main

import "fmt"

const (
	UsdToEur = 0.85
	UsdToRub = 85.00
	EurToRub = UsdToRub / UsdToEur
)

func main() {
	fmt.Printf("1 USD = %.2f EUR\n", UsdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", UsdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n", EurToRub)
}
func convert(sum float64, fromV, toV string) float64 {
	return 0
}

func readUsersInput() float64 {
	var input float64
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0
	}
	return input
}
