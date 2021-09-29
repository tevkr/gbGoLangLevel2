// Package primeNumbers provides a function that writes out all prime numbers up to the passed argument
//
// The Print(n int) writes out all prime numbers up to the number passed in the function argument
//
//  Print(n int)
//
// So, you are allowed to call this function and get some prime numbers
package primeNumbers

import "fmt"

// Print writes out all prime numbers up to the number passed in the function argument
func Print(n int) {
	var prime bool
	for i := 2; i <= n; i++ {
		prime = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Printf("%d ", i)
		}
	}
}
