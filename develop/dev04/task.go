package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
----
Полагаем что под анаграммой подразумевается множество символов в котором более одного символа
Множества анограм объединены классом эквивалентности -- лексикографический отсортированные слова анограммы совпадают посимвольно
----
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.




Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

go vet

ok

golangci-lint run --enable golint task.go

ok

*/

import (
	"fmt"
	"sort"
	"strings"
)

func sortOnAlphabet(input string) string {
	b := []byte(input)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetSetOfAnnagrams(input []string) map[string][]string {
	var tmp map[string][]string = make(map[string][]string)
	for _, v := range input {
		if len(v) > 1 {
			class := sortOnAlphabet(strings.ToLower(v))
			set, ok := tmp[class]
			if ok {
				if !contains(set, strings.ToLower(v)) {
					set = append(set, strings.ToLower(v))
					tmp[class] = set
				}
			} else {
				tmp[class] = []string{strings.ToLower(v)}
			}
		}
	}
	var ans map[string][]string = make(map[string][]string)
	for _, v := range input {
		set, ok := tmp[sortOnAlphabet(strings.ToLower(v))]
		if ok && len(set) > 1 {
			t := set[0]
			sort.Strings(set)
			ans[t] = set
		}
	}
	return ans
}

func main() {
	v := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(GetSetOfAnnagrams(v))
}
