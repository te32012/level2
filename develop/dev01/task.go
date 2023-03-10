package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.

go vet без ошибок
golangci-lint run --enable golint task.go
ok
*/

func main() {

	logger := log.New(os.Stderr, "", 0)

	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err.Error())
	}
	fmt.Println(t)
	response, er := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if er != nil {
		logger.Fatal(err.Error())
	}
	timeonserver := time.Now().Add(response.ClockOffset)
	fmt.Println(timeonserver)

}
