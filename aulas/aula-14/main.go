package main

import (
	"sync"
	"strconv"
	"time"
	"fmt"
)

var wg sync.WaitGroup

func hardTask(name string){
	defer wg.Done() 
	//Quando finalizar, a função Done será chamada, dando término ao wait
	for i := 1; i <= 10; i++ {
		time.Sleep(1 + time.Second)
		fmt.Printf("Hard Task %s: %d\n", name, i)
	}
	fmt.Printf("Hard Task [DONE]\n")
}

func main(){
	for i := 1; i <= 10; i++{
		wg.Add(1)
		go hardTask(strconv.Itoa(i))
	}
	wg.Wait()

	fmt.Println("All done\n")
}
