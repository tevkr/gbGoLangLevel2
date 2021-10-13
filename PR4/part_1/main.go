package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var workers = make(chan *int, 50)
	x := 0
	for i := 0; i < 1000; i++ {
		workers <- &x
		go func(){
			x := <-workers
			*x++
		}()
	}
	time.Sleep(time.Second)
	close(workers)
	fmt.Println("x = " + strconv.Itoa(x))
}