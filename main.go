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
