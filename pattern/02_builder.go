package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Паттерн Строитель  используется, когда нужный продукт сложный и требует нескольких шагов для построения.
В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.
При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного нестабильного продукта.
Паттерн "Строитель" скрывает объект до тех пор, пока он не построен до конца.

*/

type unfinishedHouse interface {
	buildARoof() unfinishedHouse
	buildAwindow() unfinishedHouse
	buildallhouse() finishedhouse
}

type finishedhouse interface {
	liveInHouse() string
}

type unfinishedwoodhouse struct {
	nameHouse finishedwoodhouse
}

type finishedwoodhouse struct {
	peopleliveinthishouse string
}

func (house *finishedwoodhouse) liveInHouse() string {
	return house.peopleliveinthishouse
}

func (wh *unfinishedwoodhouse) buildARoof() unfinishedwoodhouse {
	wh.nameHouse.peopleliveinthishouse += " with roof"
	return *wh
}
func (wh *unfinishedwoodhouse) buildAwindow() unfinishedwoodhouse {
	wh.nameHouse.peopleliveinthishouse += " with window"
	return *wh
}
func (wh *unfinishedwoodhouse) buildallhouse() finishedhouse {
	wh.nameHouse.peopleliveinthishouse = "people live in house " + wh.nameHouse.peopleliveinthishouse
	return &wh.nameHouse
}
