package main

import (
	"fmt"
	"time"
)

func worker(valores chan int, resultados chan int) {

	for valor := range valores {
		fmt.Printf("Recebi esse valor %d\n", valor)

		time.Sleep(time.Minute)
		resultados <- valor * 10
		fmt.Printf("Devolvendo o valor depois de %s", time.Now().Format("15:05:04"))
	}
}

func main() {
	valores := make(chan int)
	resultados := make(chan int)

	go worker(valores, resultados)
	valores <- 10

	res := <-resultados

	fmt.Println(res)
	close(valores)
	close(resultados)
}
