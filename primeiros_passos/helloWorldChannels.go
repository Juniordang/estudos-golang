package main

import (
	"fmt"
	"sync"
	"time"
)

// responsável por gerir as pausas do goroutine
var wg sync.WaitGroup

func printer(ch chan int) {
	fmt.Println("Entrei na função printer")

	for i := range ch {
		fmt.Println("Tamanho do canal: ", ch) // não consigo vê o tamanho

		fmt.Printf("Recebi o valor %d\n", i)

		time.Sleep(5 * time.Second) // faz uma pausa de 5 segundos
		fmt.Printf("o %d foi processado depois de %s\n", i, time.Now().Format("15:04:05"))
	}

	wg.Done()
}

func main() {
	ch := make(chan int)
	go printer(ch)
	wg.Add(1)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Tentando enviar o valor %d\n", i)
		ch <- i
	}
	fmt.Println("Saí do loop")

	close(ch)
	wg.Wait()
	fmt.Println("Depois do wait") // só pode ser executudo quando a gorountine terminar.
}
