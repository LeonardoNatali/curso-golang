package main

import(
	"time"
	"fmt"
)

func main(){
	done1 := make(chan bool)
	go task(done1)
	//Bloqueia o fluxo da execução até que alguma informação seja passada 
	<- done1 //Faz a leitura das informações no channel

	msg := make(chan string)
	done := make(chan bool)

	go sendMessage(msg, done)
	go receiveMessage(msg, done)
	
	<-done
}

func task(done chan bool){
	fmt.Println("START TASK")
	time.Sleep(time.Second*3)
	fmt.Println("END TASK")
	done <- true
}

func sendMessage(msg chan string, done chan bool){
	msg <- "Mensagem enviada"
}
func receiveMessage(msg chan string, done chan bool){
	fmt.Printf("Mensagem recebida: [%s]", <-msg)
	done <- true
}