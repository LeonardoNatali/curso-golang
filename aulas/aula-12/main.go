package main

import (
	"fmt"
)

type User struct {
	name string
	username string
	online bool
}

type Admin struct {
	User
	manager bool
}

type NormalUser User

func (u *NormalUser) ShowMessage(){
	u.name = "[NORMAL USER] " + u.name 
	fmt.Printf("Hello, my normal name is %s\n", u.name)
}

func (u *User) ShowUserInfo(){
	fmt.Printf("Hello, my name is %s\n", u.name)
}

func main(){
	var adm Admin
	adm.name     = "Leonardo Natali"
	adm.online   = true
	adm.username = "leonardo"
	adm.manager  = true
	adm.ShowUserInfo()
	var normal NormalUser
	normal.name = "Leonardo Natali"
	normal.ShowMessage()
}