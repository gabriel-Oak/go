package main

import (
	"fmt"
	"runtime"
	"time"
)

func fale(pessoa, texto string, qtde int, finished *bool) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
		runtime.Gosched() // libera o processador para outras goroutines
	}
	*finished = true
}

func main() {
	finishedMaria, finishedJoao := false, false

	runtime.GOMAXPROCS(1) // quantidade de processadores que a aplicação pode usar
	go fale("Maria", "Pq vc não fala comigo?", 60, &finishedMaria)
	go fale("João", "Só posso falar depois de você!", 60, &finishedJoao)

	for (finishedMaria == false) || (finishedJoao == false) {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
