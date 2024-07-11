package main

import (
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	birthdate string
	created   time.Time
}

func Struckt() {
	userfirstname := getuserdetail("enter firstname:")
	userlastname := getuserdetail("enter lastname:")
	userbirthdate := getuserdetail("enter birthdate:")
	appuser := new(userfirstname, userlastname, userbirthdate)

	appuser.outputstruct()
	appuser.clearstruct()
	appuser.outputstruct()
}
func (n *User) outputstruct() {
	fmt.Println(n.firstname, n.lastname, n.birthdate, n.created)
}
func (n *User) clearstruct() {
	n.firstname = ""
	n.lastname = ""
}
func new(userfirstname, userlastname, userbirthdate string) *User {
	return &User{
		firstname: userfirstname,
		lastname:  userlastname,
		birthdate: userbirthdate,
		created:   time.Now(),
	}
}
func getuserdetail(inputtxt string) string {
	var info string
	fmt.Println(inputtxt)
	fmt.Scan(&info)
	return info
}
