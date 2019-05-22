package main

import (
	"fmt"
	"time"
)

func main() {

	var u username
	u = "Leonardo Natali"
	fmt.Println(u)

	usuario := user{
		birthdate: time.Now().AddDate(-10, 2, 2),
		name:      "Leonardo Natali",
		live:      true,
	}
	fmt.Println(usuario.getAge())
}

type user struct {
	name      string
	birthdate time.Time
	live      bool
}

type username string

func (u *user) getAge() float64 {
	return time.Now().Sub(u.birthdate).time
}
