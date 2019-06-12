package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Procesando")

	ch := make(chan int, 2)

	// go routine lanza funcion anonima
	// cgr: control go routine
	go func() {
		values := make([]int, 0)
		for i := 0; i < 2; i++ {
			v := <-ch
			values = append(values, v)
		}
		var result int
		for _, v := range values {
			result += v
		}
		fmt.Println("Resultado", result)
		fmt.Println("Tiempo", time.Since(start))
	}()

	go computeA(ch)
	go computeB(ch)
}

func computeA(ch chan int) {
	time.Sleep(
		5 * time.Second,
	)
	fmt.Println("Retorno valor", 5)
	ch <- 5
}

func computeB(ch chan int) {
	time.Sleep(
		3 * time.Second,
	)
	fmt.Println("Retorno valor", 6)
	ch <- 6
}
