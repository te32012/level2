package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func MultiCannalUnion(channals ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(channals))
	for _, channal := range channals {
		go func(ch <-chan interface{}) {
			for element := range ch {
				out <- element
			}
			waitGroup.Done()
		}(channal)
	}
	go func() {
		waitGroup.Wait()
		close(out)
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	t := time.Now()
	<-MultiCannalUnion(
		sig(6*time.Second),
		sig(5*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(1*time.Second),
	)
	fmt.Println("прошло времени с начала запуска программы", time.Since(t))
}
