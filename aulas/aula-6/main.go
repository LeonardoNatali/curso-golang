package main

import "fmt"

func main() {
	slice := make([]int, 4, 5)
	slice[0] = 0
	slice[1] = 1
	slice[2] = 2
	fmt.Println(slice)

	slice2 := slice[1:] //Pega da 2ª posição em diante (Inclusive)
	slice2 = slice[:1]  //Pega da 2ª posição para trás
	slice2 = slice[1:2] //Pega exclusivamente o valor da 2ª posição
	slice2 = slice[:]   //Pega todos os elementos do slice de origem

	slice2 = append(slice2, 3)

	fmt.Println(slice2)
}
