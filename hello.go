package main

import (
	"example/hello/note"
	"fmt"
)

var pchoice int

func main() {

	fmt.Print("enter choice:\n")
	fmt.Print("1.profit calculator\n2.bank\n3.investment\n4>struct function")
	fmt.Scan(&pchoice)
	switch pchoice {
	case 1:
		//profit calculator
		Revenue()
		//bankdumy
	case 2:
		Bankoperation()
		//investment
	case 3:
		Investment1()
	case 4:
		Struckt()
	case 5:
		noteinfo()
	}
}
func noteinfo() {
	title := getuserdetail("enter title:")
	content := getuserdetail("enter content:")
	note.New(title, content)

}
