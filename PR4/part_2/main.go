package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChanel := make(chan os.Signal, 1)
	exitChannel := make(chan int)
	signal.Notify(signalChanel, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <- signalChanel:
				fmt.Println("Signal terminte triggered.")
				exitChannel <- 0
			case <- time.After(1 * time.Second):
				fmt.Println("Timed out.")
				exitChannel <- 1
			}
		}
	}()
	exitCode := <-exitChannel
	os.Exit(exitCode)
}