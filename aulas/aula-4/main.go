package main
import "fmt"
var res int
var msg string = ""

func main(){
	res, msg = calc(4)
	fmt.Println(res, msg)
}

func calc(n int) (int, string){
	const MAX = 100000
	acc := 0
	msg := ""
	i := 0

	for i = n; i > 1; i-- {
		acc += (i * (acc+i))
	}

	if acc >= MAX {
		msg = fmt.Sprintf("Resultado é maior ou igual a %d", MAX)
	}else{
		msg = fmt.Sprintf("Resultado é menor que %d", MAX)
	}

	return acc, msg
}