package main

import (
	"fmt"
	"time"
)

/*
	Дополните функцию из п.1 возвратом собственной ошибки в случае
	возникновения панической ситуации. Собственная ошибка должна хранить
	время обнаружения панической ситуации. Критерием успешного выполнения
	задания является наличие обработки созданной ошибки в функции main и
	вывод ее состояния в консоль.
*/

type MyError struct {
	text string
	errTime time.Time
}

func New(text string) error {
	return &MyError{
		text: text,
		errTime: time.Now(),
	}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s\ntime: %s", e.text, e.errTime)
}


func divideByZero() (err error) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
			err = New(fmt.Sprint(v))
		}
	}()
	x := 0
	y := 5 / x; _ = y
	return err
}

func main() {
	err := divideByZero()
	if err != nil {
		fmt.Println(err)
	}
}