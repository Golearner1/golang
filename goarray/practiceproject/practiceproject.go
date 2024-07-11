package practiceproject

import "fmt"

func practiceproject() {
	price := []float64{10, 20, 30, 40}
	taxrate := []float64{0, 0.07, 0.1, 0.15}
	result := make(map[float64][]float64)
	for _, taxrate := range taxrate {
		profitamount := make([]float64, len(price))
		for pindex, price := range price {
			profitamount[pindex] = price * (1 + taxrate)
		}
		result[taxrate] = profitamount
	}
	fmt.Println(result)
}
