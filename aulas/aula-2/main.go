package main
import "fmt"

func main(){
	var b = true

	for i := 1; i <= 100; i++{
		b = i % 2 == 0
		if b {
			fmt.Println("i: ", i, " - Par");
		}else{
			fmt.Println("i: ", i, " - Ímpar")
		}
	}
}