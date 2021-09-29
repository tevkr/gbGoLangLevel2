package randomStringPrinter

import (
	"fmt"
	_ "github.com/valyala/fasthttp"
	"math/rand"
	"time"
)

func SetRandSeedTimeNow() {
	rand.Seed(time.Now().UnixNano())
}

func Print(length int) {
	fmt.Print(randStringBytes(length))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}