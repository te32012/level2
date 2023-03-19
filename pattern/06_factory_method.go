package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Используйте паттерн фабричный метод, когда:
-- классу заранее неизвестно, объекты каких классов ему нужно создавать;

-- класс спроектирован так, чтобы объекты, которые он создает, специфици-
ровались подклассами;

-- класс делегирует свои обязанности одному из нескольких вспомогательных
подклассов, и вы планируете локализовать знание о том, какой класс при-
нимает эти обязанности на себя.

плюсы
-- Избавляет класс от привязки к конкретным классам продуктов.
-- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
-- Упрощает добавление новых продуктов в программу.
-- Реализует принцип открытости/закрытости.

минусы
--
Может привести к созданию больших параллельных иерархий классов,
так как для каждого класса продукта надо создать свой подкласс создателя.

*/

type Person interface {
	createContent(string) Content
}

type Content interface {
	sell()
}

type SomePerson struct {
}

func (cf *SomePerson) createContent(str string) Content {
	switch {
	case str == "video":
		return &Video{}
	case str == "music":
		return &Music{}
	case str == "txt":
		return &Text{}
	default:
		return nil
	}
}

type Music struct {
}

func (m *Music) sell() {

}

type Video struct {
}

func (v *Video) sell() {

}

type Text struct {
}

func (t *Text) sell() {
	var person Person
	var content Content = person.createContent("video")
}
