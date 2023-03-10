package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
количество совпадений с шаблоном
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
выводятся те строки которые не соответсвуют шаблону
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
дополнительно в начале распечатывать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

go vet

ok

golangci-lint run --enable golint task.go

ok

go test

ok

*/

type ConfigGrep struct {
	after     int
	before    int
	context   int
	count     bool
	ignore    bool
	invert    bool
	fixed     bool
	line      int
	inputfile string
	pattern   string
}

func (c *ConfigGrep) ConfigGrep() {
	flag.IntVar(&c.after, "A", 0, "N строк после совпадения")
	flag.IntVar(&c.before, "B", 0, "N строк после совпадения")
	flag.IntVar(&c.context, "C", 0, "+-N строк возле совпадения")
	count := flag.Bool("c", false, "Подсчет количества строк - совпадений")
	ignore := flag.Bool("i", false, "Игнорироовать регистр")
	invert := flag.Bool("v", false, "Вместо множества строк сопадений печать множества строк не вхождения")
	fixed := flag.Bool("F", false, "Точное совпадение co строкой")
	flag.IntVar(&c.line, "n", 0, "Печать линии под номером N")

	flag.Parse()

	c.count = *count
	c.ignore = *ignore
	c.invert = *invert
	c.fixed = *fixed

	arg := flag.Args()
	if len(arg) != 2 {
		fmt.Printf("В функцию был передано %v аргументов вместо 2", len(arg))
		os.Exit(1)
	}
	c.inputfile = arg[1]
	c.pattern = arg[0]
}

func (c *ConfigGrep) ReadRowsFromFile() ([]string, error) {
	file, err := os.Open(c.inputfile)
	if err != nil {
		log.Println("Ошибка чтения файла : " + err.Error())
		return nil, err
	}
	defer file.Close()

	var ans []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ans = append(ans, scanner.Text())
	}
	return ans, nil
}

func (c *ConfigGrep) Grep(s []string) (string, error) {
	if c.ignore {
		c.pattern = "(?i)" + c.pattern
	}
	re1, err := regexp.Compile(c.pattern)
	if err != nil {
		return "", err
	}

	var tmperror error = errors.New("временная ошибка")
	var sb strings.Builder

	switch {
	case c.count:
		var count int = 0
		for _, x := range s {
			if !c.fixed {
				if re1.MatchString(x) {
					if !c.invert {
						count++
					}
				} else {
					if c.invert {
						count++
					}
				}
			} else {
				if containsString(x, c.pattern) {
					if !c.invert {
						count++
					}
				} else {
					if c.invert {
						count++
					}
				}

			}
		}
		sb.WriteString(strconv.Itoa(count) + "\n")
	case c.line > 0:
		if c.line < len(s) {
			sb.WriteString(s[c.line])
		} else {
			tmperror = errors.New("количество строк в файле меньше индекса указанной строки")
		}
	case c.after > 0:
		for index, str := range s {
			if !c.fixed {
				if re1.MatchString(str) {
					if !c.invert {
						for x := index; x < index+c.after+1; x++ {
							if x >= 0 && len(s) > x {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				} else {
					if c.invert {
						for x := index; x < index+c.after+1; x++ {
							if x >= 0 && re1.MatchString(s[x]) {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				}
			} else {
				if containsString(str, c.pattern) {
					if !c.invert {
						for x := index; x < index+c.after+1; x++ {
							if x >= 0 && len(s) > x {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				} else {
					if c.invert {
						for x := index; x < index+c.after+1; x++ {
							if x >= 0 && re1.MatchString(s[x]) {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				}
			}
		}
	case c.before > 0:
		for index, str := range s {
			if !c.fixed {
				if re1.MatchString(str) {
					if !c.invert {
						for x := index - c.before; x < index+1; x++ {
							if x >= 0 && len(s) > x {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				} else {
					if c.invert {
						for x := index - c.before; x < index+1; x++ {
							if x >= 0 && re1.MatchString(s[x]) {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				}
			} else {
				if containsString(str, c.pattern) {
					if !c.invert {
						for x := index - c.before; x < index+1; x++ {
							if x >= 0 && len(s) > x {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				} else {
					if c.invert {
						for x := index - c.before; x < index+1; x++ {
							if x >= 0 && re1.MatchString(s[x]) {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				}

			}
		}
	case c.context > 0:
		for index, str := range s {
			if !c.fixed {
				if re1.MatchString(str) {
					if !c.invert {
						for x := index - c.context; x < index+c.context+1; x++ {
							if x >= 0 && len(s) > x {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				} else {
					if c.invert {
						for x := index - c.context; x < index+c.context+1; x++ {
							if x >= 0 && re1.MatchString(s[x]) {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				}
			} else {
				if containsString(str, c.pattern) {
					if !c.invert {
						for x := index - c.context; x < index+c.context+1; x++ {
							if x >= 0 && len(s) > x {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				} else {
					if c.invert {
						for x := index - c.context; x < index+c.context+1; x++ {
							if x >= 0 && re1.MatchString(s[x]) {
								sb.WriteString(s[x] + "\n")
							}
						}
					}
				}
			}
		}
	default:
		for _, str := range s {
			if !c.fixed {
				if re1.MatchString(str) {
					if !c.invert {
						sb.WriteString(str + "\n")
					}
				} else {
					if c.invert {
						sb.WriteString(str + "\n")
					}
				}
			} else {
				if containsString(str, c.pattern) {
					if !c.invert {
						sb.WriteString(str + "\n")
					}
				} else {
					if c.invert {
						sb.WriteString(str + "\n")
					}
				}
			}
		}
	}
	if tmperror.Error() != "временная ошибка" {
		return sb.String(), tmperror
	}
	return sb.String(), nil
}

func containsString(base string, pattern string) bool {
	var cont bool
	for x := 0; x < len(base)-len(pattern)+1; x++ {
		cont = true
		for y := 0; y < len(pattern); y++ {
			if base[x+y] != pattern[y] {
				cont = false
				break
			}
		}
		if cont {
			break
		}
	}
	return cont
}

func main() {

	var conf ConfigGrep
	conf.ConfigGrep()
	rows, err := conf.ReadRowsFromFile()
	if err != nil {
		log.Fatal(err)
	}
	ans, err := conf.Grep(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(ans)
}
