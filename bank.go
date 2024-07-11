package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const textfile = "balance.txt"

func Bankoperation() {

	fmt.Println("bankdumy")

	investmentamount, errorr := Readbalance()
	if errorr != nil {
		fmt.Println("error :", errorr)
		fmt.Printf("----------")
		panic("an unknown error occured")
	}
	var choice, withdraw float64
	fmt.Println("welcome to the bank\nenteer your choice\n1.deposit\n2.withdraw\n3.check balance\n4.exit")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("enter the amount to invest")
		fmt.Scan(&investmentamount)
		Writebalance(investmentamount)
		if investmentamount <= 0 {
			fmt.Println("this amount can not be added")
			break
		}
		investmentamount = +investmentamount
		Writebalance(investmentamount)
		fmt.Println("your account balance is", investmentamount)

	case 2:
		fmt.Println("enter amount to withdraw")
		fmt.Scan(&withdraw)
		if withdraw > investmentamount {
			fmt.Println("insufficient balance")
			break
		}
		investmentamount = investmentamount - withdraw
		Writebalance(investmentamount)
		fmt.Println("amount remining", investmentamount)

	case 3:
		fmt.Println("you balance is :", investmentamount)

	default:
		fmt.Println("thanks for visiting bank")
	}
}
func Writebalance(balance float64) {
	balancetext := fmt.Sprint(balance)
	os.WriteFile(textfile, []byte(balancetext), 0644)
}
func Readbalance() (float64, error) {
	data, err := os.ReadFile(textfile)
	if err != nil {
		return 10000, errors.New("read failed")
	}
	balancetext := string(data)
	balance, err := strconv.ParseFloat(balancetext, 64)
	if err != nil {
		return 10000, errors.New("conversion failed")
	}
	return balance, nil
}
