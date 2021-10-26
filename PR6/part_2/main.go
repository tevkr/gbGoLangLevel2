package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

/*
	Написать многопоточную программу, в которой будет использоваться
	явный вызов планировщика. Выполните трассировку программы

	go tool trace trace.out
*/

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	go log.Println("I'm working!")
	for i := 0;i <= 1e9; i += 1 {
		if i%1e6 == 0 {
			runtime.Gosched()
		}
	}
}
