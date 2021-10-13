package main

import (
	"fmt"
	"sync"
)

/*
	Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”

	go run -race main.go

	WARNING: DATA RACE
*/

const count = 1000

func main() {
	var (
		counter int
		wg sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
