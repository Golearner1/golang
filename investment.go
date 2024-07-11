package main

import (
	"fmt"
	"math"
)

func Investment1() {
	fmt.Println("investment")
	var investmentt, futurevalue, years, amount, intrest float64
	fmt.Println("enter amunt you want to invest")
	fmt.Scan(&amount)
	fmt.Println("for how much year u want to invest")
	fmt.Scan(&years)
	fmt.Println("expexted return per year")
	fmt.Scan(&intrest)
	investmentt = amount * years
	fmt.Println("total investmet", investmentt)
	futurevalue = investmentt / math.Pow(1+intrest/100, years)
	fmt.Println("real value after investement is made constantly", futurevalue)
}
