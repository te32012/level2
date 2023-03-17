package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

плюсы
-- Избавляет от множества больших условных операторов машины состояний.
-- Концентрирует в одном месте код, связанный с определённым состоянием.
-- Упрощает код контекста.

минусы
-- Может неоправданно усложнить код, если состояний мало и они редко меняются.

*/

type PartMessage interface {
	DecodeAndUseMessage() string
}

type HeaderPartMessage struct {
}

func (hpm *HeaderPartMessage) DecodeAndUseMessage() string {
	return "use 107FM"
}

type MainPartMessage struct {
}

func (mpm *MainPartMessage) DecodeAndUseMessage() string {
	return "Use 110FM"
}

type EndPartMessage struct {
}

func (epm *EndPartMessage) DecodeAndUseMessage() string {
	return "Use 100FM"

}

type Receiver struct {
	message PartMessage
}

func (r *Receiver) Receiver() {
	r.message = &HeaderPartMessage{}
}
func (r *Receiver) SetPartMessage(msg PartMessage) {
	r.message = msg
}
