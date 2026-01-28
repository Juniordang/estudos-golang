package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		time.Sleep(5 * time.Second)
		fmt.Printf("Recebi o %d as %s\n", i, time.Now().Format("15:04:07"))
	}

	wg.Done()
}

func main() {
	ch := make(chan int)
	go printer(ch)
	wg.Add(1)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Tentando enviar o %d \n", i)
		ch <- i
		fmt.Printf("Envio do %d confirmado \n", i)
	}

	close(ch)
	wg.Wait()
}
