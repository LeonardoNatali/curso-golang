package main

import "fmt"

func main() {
	var l List
	l.Init()
	l.Show()
}

type List []interface{}

func (l *List) Init() {
	*l = []interface{}{
		10,
		"Olá",
		1.9,
		true,
	}
}

func (l List) Show() {
	fmt.Println(l...)
}
