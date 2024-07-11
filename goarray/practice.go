package main

import (
	"example.com/first-app/price"
)

func main() {
	taxrate := []float64{0, 0.07, 0.1, 0.15}
	for _, taxrate := range taxrate {
		taxprice := price.New(taxrate)
		taxprice.Process()
	}
}
