package main

import (
	"fmt"
	"sync"
)

/*
	Напишите программу, которая запускает n потоков и дожидается завершения их всех
*/

func main() {
	var (
		wg = sync.WaitGroup{}
		n int
	)
	fmt.Print("n = ")
	fmt.Scan(&n)
	wg.Add(n)
	for i := 0; i < n; i += 1 { // Запускаем n потоков
		go func() {
			wg.Done()
		}()
	}
	wg.Wait() // Дожидаемся
}
