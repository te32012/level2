package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.

go vet

ok

golangci-lint run --enable golint task.go

ok

go test

ok

*/

func main() {
	str := "\\45a55"
	s, e := UnpacStr(str)
	fmt.Println(s)
	fmt.Println(e)
}

func UnpacStr(str string) (string, error) {
	var sb strings.Builder = strings.Builder{}
	var comperessedSymbol rune = 0
	var numberOfCompressedSymbol int
	var isEscapeSymbol bool

	for _, symbol := range str {
		switch {
		case (symbol == '\\') && !isEscapeSymbol:
			if comperessedSymbol != 0 && numberOfCompressedSymbol > 0 {
				for x := 0; x < numberOfCompressedSymbol; x++ {
					_, e := sb.WriteRune(comperessedSymbol)
					if e != nil {
						return "", e
					}
				}
			} else {
				if comperessedSymbol != 0 {
					_, e := sb.WriteRune(comperessedSymbol)
					if e != nil {
						return "", e
					}
				}
			}
			comperessedSymbol = 0
			numberOfCompressedSymbol = 0
			isEscapeSymbol = true

		case (symbol >= '0' && symbol <= '9') && !isEscapeSymbol:
			numberOfCompressedSymbol = (numberOfCompressedSymbol * 10) + int(symbol-'0')
		case (symbol >= '0' && symbol <= '9') && isEscapeSymbol:
			comperessedSymbol = symbol
			isEscapeSymbol = false
		default:
			if comperessedSymbol != 0 && numberOfCompressedSymbol > 0 {
				for x := 0; x < numberOfCompressedSymbol; x++ {
					_, e := sb.WriteRune(comperessedSymbol)
					if e != nil {
						return "", e
					}

				}
			} else {
				if comperessedSymbol != 0 {
					_, e := sb.WriteRune(comperessedSymbol)
					if e != nil {
						return "", e
					}
				} else {
					if numberOfCompressedSymbol > 0 {
						return "", errors.New("uncorect string")
					}
				}
			}
			numberOfCompressedSymbol = 0
			comperessedSymbol = symbol
			isEscapeSymbol = false
		}
	}
	if comperessedSymbol != 0 && numberOfCompressedSymbol > 0 {
		for x := 0; x < numberOfCompressedSymbol; x++ {
			_, e := sb.WriteRune(comperessedSymbol)
			if e != nil {
				return "", e
			}
		}
	} else {
		if comperessedSymbol != 0 {
			_, e := sb.WriteRune(comperessedSymbol)
			if e != nil {
				return "", e
			}
		} else {
			if numberOfCompressedSymbol > 0 {
				return "", errors.New("uncorect string")
			}
			return "", nil

		}
	}

	return sb.String(), nil
}
