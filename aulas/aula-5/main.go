//Array, slice e Map
package main
import "fmt"

func main(){
	const VALOR = 256
	i := 0
	
	//Array
	itens := [10]int{}
	for i = 0; i < len(itens); i++{
		itens[i] = ((i+1) * VALOR)
	}
	fmt.Println(itens)
	
	//Map
	values := map[string] string{
		"name":"Leonardo",
		"last_name": "Natali",
	}

	values ["name"] = "LEONARDO"
	fmt.Println(values)
	fmt.Println("Sobrenome:", values["last_name"])
	fmt.Println("Tamanho do map:", len(values))

	//Slice
	slice := [] int {0}
	for i = 0; i < VALOR; i++ {
		slice = append(slice, (i * VALOR))
	}

	fmt.Println(slice)
}