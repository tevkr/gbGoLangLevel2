package main

import (
	"fmt"
)

/*
	Для закрепления навыков отложенного вызова функций, напишите программу,
	содержащую вызов функции, которая будет создавать паническую ситуацию неявно.
	Затем создайте отложенный вызов, который будет обрабатывать эту паническую ситуацию
	и, в частности, печатать предупреждение в консоль. Критерием успешного выполнения
	задания является то, что программа не завершается аварийно ни при каких условиях.
*/

func divideByZero() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	x := 0
	y := 5 / x; _ = y
}

func main() {
	divideByZero()
}