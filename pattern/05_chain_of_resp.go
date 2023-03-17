package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Используйте цепочку обязанностей, когда:

--  есть более одного объекта, способного обработать запрос, причем настоя-
щий обработчик заранее неизвестен и должен быть найден автоматически;

-- вы хотите отправить запрос одному из нескольких объектов, не указывая
явно, какому именно;

-- набор объектов, способных обработать запрос, должен задаваться динами-
чески.

плюсы
-- Уменьшает зависимость между клиентом и обработчиками.
-- Реализует принцип единственной обязанности.
-- Реализует принцип открытости/закрытости.

минусы
-- Запрос может остаться никем не обработанным.

*/

type HandlerHttpResponse interface {
	getResponse(int)
}
type HandlerPageNotFound struct {
	next HandlerHttpResponse
}

func (hpnf *HandlerPageNotFound) getResponse(status int) {
	if status == 404 {
		fmt.Println("страница не найдена")
	} else if hpnf.next != nil {
		hpnf.getResponse(status)
	}
}

type HandlerOk struct {
	next HandlerHttpResponse
}

func (hpnf *HandlerOk) getResponse(status int) {
	if status == 200 {
		fmt.Println("ok")
	} else if hpnf.next != nil {
		hpnf.getResponse(status)
	}

}
