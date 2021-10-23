package main

import (
	"fmt"
	"sync"
)

/*
	Реализуйте функцию для разблокировки мьютекса с помощью defer
*/

// Функция для разблокировки мьютекса с помощью defer
func criticalSection(mtx *sync.Mutex) {
	defer mtx.Unlock()
}

func main() {
	var mtx sync.Mutex
	mtx.Lock()
	criticalSection(&mtx)
	mtx.Lock()
	criticalSection(&mtx)
	fmt.Print("done")
}


