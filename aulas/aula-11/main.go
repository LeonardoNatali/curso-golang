package main

import (
	"fmt"
)

type User struct {
	name string
	online bool
}

type Admin struct{
	User
	ranking int
}

func(u *User) Login(){
	u.online = true;
	fmt.Printf("\nO usuário %s está logado\n", u.name)
}

func(u *User) Logout(){
	u.online = false;
	fmt.Printf("\nO usuário %s saiu do sistema\n", u.name)
}

type NormalUser User

func (u *NormalUser) ShowMessage(){
	u.name = "[NORMAL USER] " + u.name 
	fmt.Printf("Hello, my normal name is %s\n", u.name)
}

func main(){
	var adm  Admin
	adm.ranking = 2
	adm.name = "Leonardo"
	adm.Login()
	adm.Logout()
	fmt.Println(adm.ranking)

	var normal NormalUser
	normal.name = "Leonardo Natali"
}

