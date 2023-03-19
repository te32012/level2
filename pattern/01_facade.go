package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

	применимость

Когда вам нужно представить простой или урезанный интерфейс к сложной подсистеме.
Когда вы хотите разложить подсистему на отдельные слои.
+
Изолирует клиентов от компонентов сложной подсистемы.
-
Фасад рискует стать божественным объектом, привязанным ко всем классам программы.

*/

// example
type car struct {
	eng   *engine
	stwh  *steeringWheel
	cdash *carDashboard
}

func (c *car) use() string {
	return c.cdash.lookatcarDashboard() + c.stwh.usesteeringWheel() + c.eng.startengine()
}

type engine struct {
}

func (e *engine) startengine() string {
	return "start engine"
}

type steeringWheel struct {
}

func (sw *steeringWheel) usesteeringWheel() string {
	return "use steeringWheel"
}

type carDashboard struct {
}

func (cd *carDashboard) lookatcarDashboard() string {
	return "look at car dashboard"
}

func main() {
	var c car
	fmt.Println(c.cdash.lookatcarDashboard())
	fmt.Println(c.eng.startengine())
	fmt.Println(c.stwh.usesteeringWheel())
}
