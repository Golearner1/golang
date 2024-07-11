package recurssion

import "fmt"

func recurssion1() {
	result := recurssion(5)
	fmt.Println(result)
}
func recurssion(number int) int {
	if number == 0 {
		return 1
	}
	return number * recurssion(number-1)
}
