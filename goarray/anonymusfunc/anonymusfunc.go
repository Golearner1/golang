package anonymusfunc

import "fmt"

func anonymusfunc() {
	numbers := []int{1, 2, 3, 4}
	tnumber := transformnumber(&numbers, func(val int) int {
		return val * 2
	})
	fmt.Println(tnumber)
	double := mulnumber(2)
	treiple := mulnumber(3)
	doublen := transformnumber(&numbers, double)
	triplen := transformnumber(&numbers, treiple)
	fmt.Println(doublen)
	fmt.Println(triplen)

}
func transformnumber(numbers *[]int, transform func(int) int) []int {
	dnumbers := []int{}
	for _, val := range *numbers {
		dnumbers = append(dnumbers, transform(val))
	}
	return dnumbers
}
func mulnumber(variable int) func(int) int {
	return func(val int) int {
		return val * variable
	}
}
