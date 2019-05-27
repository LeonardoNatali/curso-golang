package main

import (
	"fmt"
)

type User struct {
	name string
	username string
	online bool
}

func(u User)Show() string{
	return fmt.Sprintf("Hello, my name is %s\n", u.name)
}

func(u Admin)Show() string{
	return fmt.Sprintf("Hello, my name is %s And i am an Adm\n", u.name)
}

type Admin struct {
	User
	manager bool
}

func ShowUserInfo(u UsersInterface){
	fmt.Println(u.Show())
}

type UsersInterface interface {
	Show() string
}

func main(){
	adm := Admin{
		User{
			"Leonardo Natali",
			"leonardo",
			true,
		},
		true,
	}
	ShowUserInfo(adm)
}