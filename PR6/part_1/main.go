package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

/*
	Написать программу, которая использует мьютекс для безопасного
	доступа к данным из нескольких потоков. Выполните трассировку программы

	go tool trace trace.out
*/

const count = 1000

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

	var (
		counter int
		lock sync.Mutex
		wg sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
