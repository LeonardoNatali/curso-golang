package main

import (
	"fmt"
	"time"
)

func main() {
	usuario := user{
		birthdate: time.Now(),
		name     :      "Leonardo Natali",
		live     :      true,
	}
	usuario.getAge()
}

type user struct {
	name      string
	birthdate time.Time
	live      bool
}

func (u *user) getAge() {
	fmt.Println(u.birthdate)
}