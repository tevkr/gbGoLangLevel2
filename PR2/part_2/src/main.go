package main

import (
	"fmt"
	primeNumbers "github.com/tevkr/gbGoLangLevel2/PR2/part_2/src/primeNumbers"
)

func main() {
	var x int
	fmt.Scanln(&x)
	primeNumbers.Print(x)
}
