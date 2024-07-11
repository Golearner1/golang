package variadic

import "fmt"

func variadic() {
	numbers := []int{1, 2, 3, 4}
	result := sumup(10, 20, 30, 40)
	fmt.Println(result)
	nresult := sumup(numbers...)
	fmt.Println(nresult)
}
func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
