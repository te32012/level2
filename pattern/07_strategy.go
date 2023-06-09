package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Используйте паттерн стратегия, когда:
-- имеется много родственных классов, отличающихся только поведением.
Стратегия позволяет сконфигурировать класс, задав одно из возможных по-
ведений;

-- вам нужно иметь несколько разных вариантов алгоритма. Например, мож-
но определить два варианта алгоритма, один из которых требует больше
времени, а другой - больше памяти. Стратегии разрешается применять,
когда варианты алгоритмов реализованы в виде иерархии классов [НО87];

--- в алгоритме содержатся данные, о которых клиент не должен «знать». Ис-
пользуйте паттерн стратегия, чтобы не раскрывать сложные, специфичные для алгоритма структуры данных;
-- в классе определено много поведений, что представлено разветвленными
условными операторами. В этом случае проще перенести код из ветвей в отдельные классы стратегий.

плюсы
-- Горячая замена алгоритмов на лету.
-- Изолирует код и данные алгоритмов от остальных классов.
-- Уход от наследования к делегированию.
-- Реализует принцип открытости/закрытости.

минусы
-- Усложняет программу за счёт дополнительных классов.
-- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

*/

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

type StrategyOrderGraph interface {
	Order(Node)
}

type NLR struct {
}

func (nlr *NLR) Order(n Node) {
	fmt.Println(n.Data)
	if n.Left != nil {
		nlr.Order(*n.Left)
	}
	if n.Right != nil {
		nlr.Order(*n.Right)
	}
}

type LNR struct {
}

func (lnr *LNR) Order(n Node) {
	if n.Left != nil {
		lnr.Order(*n.Left)
	}
	fmt.Println(n.Data)
	if n.Left != nil {
		lnr.Order(*n.Left)
	}

}

type LRN struct {
}

func (lrn *LRN) Order(n Node) {
	if n.Left != nil {
		lrn.Order(*n.Left)
	}
	if n.Left != nil {
		lrn.Order(*n.Left)
	}
	fmt.Println(n.Data)
}

type Context struct {
	strategy StrategyOrderGraph
}

func (c *Context) Content(s StrategyOrderGraph) {
	c.strategy = s
}

func main() {
	var context Context = Context{}
	context.Content(&LNR{})
	var l Node = Node{Data: 5}
	var r Node = Node{Data: 2}
	var n Node = Node{Data: 11, Left: &l, Right: &r}
	context.strategy.Order(n)
	context.Content(&NLR{})
	context.strategy.Order(n)

}
