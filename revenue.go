package main

import (
	"errors"
	"fmt"
	"os"
)

func Revenue() {
	var EBT float64
	var EAT float64
	var tax float64
	var ratio float64
	fmt.Println("profit calculator")
	revenue, err2 := Getuser("revenue")
	if err2 != nil {
		fmt.Print(err2)
		panic("error occurred")
	}
	expenses, err3 := Getuser("expenses:")
	if err3 != nil {
		fmt.Print(err2)
		panic("error occurred")
	}
	taxrate, err4 := Getuser("taxrate")
	if err4 != nil {
		fmt.Print(err2)
		panic("error occurred")
	}
	EBT = revenue - expenses
	tax = EBT * taxrate / 100
	EAT = EBT - tax
	ratio = EAT / EBT
	Write(EBT, EAT, ratio)
	fmt.Printf("earning before tax%v\nearning after tax%v", EBT, EAT)
	fmt.Print("\nratio:", ratio)

}
func Write(EBT, EAT, ratio float64) {
	result := fmt.Sprintf("ebt:%.1f\n eat:%.1f\n ratio:%.2f\n", EBT, EAT, ratio)
	os.WriteFile("result.txt", []byte(result), 0644)
}
func Getuser(textt string) (float64, error) {
	var profitss float64
	fmt.Print(textt)
	fmt.Scan(&profitss)
	if profitss <= 0 {
		return 0, errors.New("error occurred")
	}
	return profitss, nil
}
