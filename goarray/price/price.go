package price

import (
	"fmt"

	"example.com/first-app/conversion"
	"example.com/first-app/fileread"
)

type Taxjob struct {
	Taxrate         float64
	Price           []float64
	Taxincludeprice map[string]string
}

func New(taxrate float64) *Taxjob {
	return &Taxjob{
		Taxrate: taxrate,
		Price:   []float64{10, 20, 30, 40},
	}
}
func (job *Taxjob) Process() {
	job.read()
	result := make(map[string]string)
	for _, price := range job.Price {
		tprice := price * (1 + job.Taxrate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", tprice)
	}
	job.Taxincludeprice = result
	fileread.Write(fmt.Sprintf("result_%.0f.json", job.Taxrate*100), job)
}
func (job *Taxjob) read() {
	lines, err := fileread.Read("prices.txt")
	if err != nil {
		fmt.Println("failed to covert string")
		fmt.Print(err)
		return
	}
	prices, err := conversion.Conversion(lines)

	if err != nil {
		fmt.Println("failed to covert string")
		fmt.Print(err)
		return
	}
	job.Price = prices
}
