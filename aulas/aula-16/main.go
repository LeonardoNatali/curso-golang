package main

import(
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup

func main(){
	done := make(chan bool)
	go hardTask(done)
	wg.Add(1)
	go loader(done)

	wg.Wait()
}

func hardTask(done chan <- bool){ //Diz que o channel só pode ser escrito 
	fmt.Println("Start task")
	time.Sleep(time.Second * 1)
	done <- true
}

func loader(done <- chan bool){
	defer wg.Done()
	i := 0
	load := []rune(`|\-/`)
	for {
		select{
		case <- done:
			fmt.Println("Done")
			return
		default:
			fmt.Print(string(load[i])+"\r")
			i++

			if i == len(load)-1 {
				i = 0
			}
		}
	}
}