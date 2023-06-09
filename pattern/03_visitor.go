package pattern

import (
	"strings"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern


Паттерн Visitor определяет операцию, выполняемую на каждом элементе из некоторой структуры.
Позволяет, не изменяя классы этих объектов, добавлять в них новые операции.
Является классической техникой для восстановления потерянной информации о типе.

Паттерн Visitor позволяет выполнить нужные действия в зависимости от типов двух объектов.
Достоинства и недостатки паттерна Посетитель

Упрощает добавление новых операций.
С помощью Посетителей легко добавлять операции, зависящие от компонентов сложных объектов.
Для определения новой операции над структурой объектов достаточно просто ввести нового Посетителя.
Напротив, если функциональность распределена по нескольким классам,
то для определения новой операции придется изменить каждый класс.

Объединяет родственные операции и отсекает те,
которые не имеют к ним отношения.
Родственное поведение не разносится по всем классам,
присутствующим в структуре объектов, оно локализовано в Посетителе.
Не связанные друг с другом функции распределяются по отдельным подклассам класса Visitor.
Это способствует упрощению как классов, определяющих элементы,
так и алгоритмов, инкапсулированных в Посетителях.
Все относящиеся к алгоритму структуры данных можно скрыть в Посетителе.

Добавление новых классов Concrete Element (в случае примера - тип нод графа) затруднено.
Каждый новый конкретный элемент требует объявления новой абстрактной операции в классе Visitor (человеке посещающем граф),
которую нужно реализовать в каждом из существующих классов ConcreteVisitor.
Иногда большинство конкретных Посетителей могут унаследовать операцию по умолчанию,
предоставляемую классом Visitor, что скорее исключение, чем правило.
Поэтому при решении вопроса о том, стоит ли использовать паттерн Посетитель,
нужно прежде всего посмотреть, что будет изменяться чаще: алгоритм,
применяемый к объектам структуры, или классы объектов, составляющих эту структуру.
*/

type Graph interface {
	printValue(Data) string
}

type Data interface {
	meetsWith() string
	sayFirstnameLastname() string
}

type Person struct {
	firstname string
	lastname  string
	age       int
}

func (person Person) meetsWith() string {
	return person.firstname + " " + person.lastname + " meets with "
}
func (person Person) sayFirstnameLastname() string {
	return person.firstname + " " + person.lastname
}

type Leaf struct {
	persons []Data
}

func (node Leaf) printValue(person Data) string {
	var sb strings.Builder
	for _, x := range node.persons {
		sb.WriteString(person.meetsWith() + x.sayFirstnameLastname())
	}
	return sb.String()
}

func main() {
	var pers Data
	pers = Person{firstname: "ivan", lastname: "ivanov", age: 20}
	var graph Graph
	var per []Data = make([]Data, 0)
	per = append(per, Person{firstname: "jane", lastname: "jane", age: 17})
	per = append(per, Person{firstname: "kate", lastname: "kate", age: 21})
	leaf := Leaf{persons: per}
	graph = leaf
	graph.printValue(pers)
}
