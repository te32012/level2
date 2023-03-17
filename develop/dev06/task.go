package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

go vet
ok

golangci-lint run --enable golint task.go
ok

go test
ok

*/

type ConfCut struct {
	filed       string
	delimiter   string
	isseparated bool
}

func (cc *ConfCut) ConfCut() {
	cc.isseparated = *flag.Bool("s", false, "выводить только токенты c разделителем")
	flag.StringVar(&cc.delimiter, "d", " ", "тип разделителя токентов")
	flag.StringVar(&cc.filed, "f", "1", "группа выводимых символов")
	flag.Parse()

}
func (cc *ConfCut) ReadFromStdin() ([]string, error) {
	var ans []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ans = append(ans, scanner.Text())
	}
	return ans, nil
}
func (cc *ConfCut) Cut(lines []string) (string, error) {
	var ans strings.Builder
	var rangeOfLines []string = strings.Split(cc.filed, ",")
	for _, y := range lines {
		for j, x := range rangeOfLines {
			tmp1 := strings.Split(x, "-")
			tmp2 := strings.Split(y, cc.delimiter)
			if !cc.isseparated {
				switch {
				default:
					return "", errors.New("Некорректный формат группы выводимых строк")
				case len(tmp1) == 2:
					first, _ := strconv.Atoi(tmp1[0])
					first--
					second, err := strconv.Atoi(tmp1[1])
					if err != nil && (err.Error() == `strconv.Atoi: parsing "": invalid syntax`) {
						second = len(tmp2) - 1
					}
					for i := first; i < second; i++ {
						if i >= 0 && i < len(tmp2) {
							if ((j == (len(rangeOfLines) - 1)) && second-1 == i) || (len(tmp2) < second && len(tmp2)-1 == i) {
								ans.WriteString(tmp2[i] + "\n")

							} else {
								ans.WriteString(tmp2[i] + cc.delimiter)
							}
						}
					}
				case len(tmp1) == 1:
					first, _ := strconv.Atoi(tmp1[0])
					first--
					if first >= 0 && first < len(tmp2) {
						if (j == (len(rangeOfLines) - 1)) || (len(tmp2)-1 == first) {
							ans.WriteString(tmp2[first] + "\n")
						} else {
							ans.WriteString(tmp2[first] + cc.delimiter)
						}
					}
				}
			} else {
				if len(tmp2) > 1 {
					switch {
					default:
						return "", errors.New("Некорректный формат группы выводимых строк")
					case len(tmp1) == 2:
						first, _ := strconv.Atoi(tmp1[0])
						first--
						second, err := strconv.Atoi(tmp1[1])
						if err != nil && (err.Error() == `strconv.Atoi: parsing "": invalid syntax`) {
							second = len(tmp2)
						}
						for i := first; i < second; i++ {
							if i >= 0 && i < len(tmp2) {
								if (j == (len(rangeOfLines) - 1)) && (second-1 == i || (len(tmp2) < second && len(tmp2)-1 == i)) {
									ans.WriteString(tmp2[i] + "\n")

								} else {
									ans.WriteString(tmp2[i] + cc.delimiter)
								}
							}
						}
					case len(tmp1) == 1:
						first, _ := strconv.Atoi(tmp1[0])
						first--
						if first >= 0 && first < len(tmp2) {
							if j == (len(rangeOfLines) - 1) {
								ans.WriteString(tmp2[first] + "\n")
							} else {
								ans.WriteString(tmp2[first] + cc.delimiter)
							}
						}
					}
				}
			}
		}
	}
	return ans.String(), nil
}

func main() {
	_, e := strconv.Atoi("")
	fmt.Println(e.Error() == `strconv.Atoi: parsing "": invalid syntax`)
}
