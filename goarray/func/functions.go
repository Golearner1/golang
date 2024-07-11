package function

import "fmt"

type tfunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	anothernumbers := []int{6, 7, 8, 9}
	doubled := transformnumber(&numbers, double)
	fmt.Println(doubled)
	tripled := transformnumber(&numbers, triple)
	fmt.Println(tripled)
	tnumbers := makethevarfunc(&numbers)
	anumbers := makethevarfunc(&anothernumbers)
	doubled = transformnumber(&numbers, tnumbers)
	tripled = transformnumber(&anothernumbers, anumbers)
	fmt.Println(tripled)
	fmt.Println(doubled)
}
func transformnumber(numbers *[]int, transform tfunc) []int {
	dnumbers := []int{}
	for _, val := range *numbers {
		dnumbers = append(dnumbers, transform(val))
	}
	return dnumbers
}
func double(val int) int {
	return val * 2
}
func triple(val int) int {
	return val * 3
}
func makethevarfunc(numbers *[]int) tfunc {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}
