package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Что делает?
Инкапсулирует запрос как объект, позволяя тем самым задавать параметры
клиентов для обработки соответствующих запросов, ставить запросы в очередь
или протоколировать их, а также поддерживать отмену операций.

Плюсы команды (что можно делать с помощью команды)
-- Параметризовать объекты выполняемым действием, как в случае с пункта-
ми меню. В процедурном языке такую параметризацию можно
выразить с помощью функции обратного вызова, то есть такой функции, ко-
торая регистрируется, чтобы быть вызванной позднее. Команды представ-
ляют собой объектно-ориентированную альтернативу функциям обратного вызова;
-- определять, ставить в очередь и выполнять запросы в разное время. Время
жизни объекта Command необязательно должно зависеть от времени жизни исходного запроса. Если получателя запроса удается реализовать так, что-
бы он не зависел от адресного пространства, то объект-команду можно пе-
редать другому процессу, который займется его выполнением;
-- поддержать отмену операций. Операция Execute объекта Command может
сохранить состояние, необходимое для отката действий, выполненных ко-
мандой. В этом случае в интерфейсе класса Command должна быть допол-
нительная операция Unexecute, которая отменяет действия, выполненные
предшествующим обращением к Execute. Выполненные команды хранятся
в списке истории. Для реализации произвольного числа уровней отмены
и повтора команд нужно обходить этот список соответственно в обратном
и прямом направлениях, вызывая при посещении каждого элемента коман-
ду Unexecute или Execute;
-- поддержать протоколирование изменений, чтобы их можно было выпол-
нить повторно после аварийной остановки системы. Дополнив интерфейс
класса Command операциями сохранения и загрузки, вы сможете вести про-
токол изменений во внешней памяти. Для восстановления после сбоя нужно
будет загрузить сохраненные команды с диска и повторно выполнить их с помощью операции Execute;
-- структурировать систему на основе высокоуровневых операций, построен-
ных из примитивных. Такая структура типична для информационных сис-
тем, поддерживающих транзакции. Транзакция инкапсулирует набор изме-
нений данных. Паттерн команда позволяет моделировать транзакции. У всех
команд есть общий интерфейс, что дает возможность работать одинаково
с любыми транзакциями. С помощью этого паттерна можно легко добавлять
в систему новые виды транзакций.

Недостатки команды
Усложняет код программы из-за введения множества дополнительных классов.
*/

type UiElement interface {
	pressUiElement() string
}

type UiMenu struct {
	hardwareDevice Device
}

func (c *UiMenu) pressUiElement() {
	c.hardwareDevice.pressKey()
}

type Device interface {
	pressKey()
}

type Mouse struct {
	someUser User
}

func (m *Mouse) pressKey() {
	m.someUser.lookAtNextUIMenu()
}

type Keyword struct {
	someUser User
}

func (k *Keyword) pressKey() {
	k.someUser.lookAtPreviousUIMenu()
}

type User interface {
	lookAtNextUIMenu()
	lookAtPreviousUIMenu()
}

type SomeMan struct {
	lookAtMonitor bool
}

func (sm *SomeMan) lookAtNextUIMenu() {
	sm.lookAtMonitor = false
}
func (sm *SomeMan) lookAtPreviousUIMenu() {
	sm.lookAtMonitor = true
}
