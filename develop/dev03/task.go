package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
имеется в виду числовое значение символов
-r — сортировать в обратном порядке
имеется в виду, что row c индексом 0 имеет n место а row с индексом n имеет место 0
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
"JAN" < "FEB" < ... < "DEC"
-b — игнорировать хвостовые (реализовано любые крайние) пробелы
_abc_ превращается в abc

-c — проверять отсортированы ли данные
выбрасывать ошибку если данные не отсортированы
-h — сортировать по числовому значению с учётом суффиксов
мало примеров
не реализовано

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

go vet

ok

golangci-lint run --enable golint task.go

ok

go test

ok

*/

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type KeyOfProgramm struct {
	column         int
	numberSymbol   bool
	reverse        bool
	unique         bool
	month          bool
	ignorespace    bool
	sorted         bool
	output         bool
	monthset       []string
	filenameinput  string
	filenameoutput string
}

func (keys *KeyOfProgramm) KeyOfProgramm() {
	numberSymbol := flag.Bool("n", false, "сортировать по числовому значению символов")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	month := flag.Bool("M", false, "сортировать по названию месяца")
	ignorespace := flag.Bool("b", true, "игнорировать любые крайние пробелы")
	sorted := flag.Bool("c", false, "проверять отсортированы ли данные")
	output := flag.Bool("o", false, "второй аргумент имя файла вывода")
	flag.IntVar(&keys.column, "k", 0, "указание колонки для сортировки")

	flag.Parse()

	keys.numberSymbol = *numberSymbol
	keys.reverse = *reverse
	keys.unique = *unique
	keys.month = *month
	keys.ignorespace = *ignorespace
	keys.sorted = *sorted
	keys.output = *output
	a := flag.Args()
	switch {
	case len(a) == 1 && !keys.output:
		keys.filenameinput = a[0]
	case len(a) == 2 && keys.output:
		keys.filenameinput = a[0]
		keys.filenameoutput = a[1]
	default:
		log.Fatal("ошибка в количестве аргументов переданных в командную строку")
	}
}

func (keys *KeyOfProgramm) ReadRowsFromFile() ([]string, error) {
	file, err := os.Open(keys.filenameinput)
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

func (keys *KeyOfProgramm) SaveRowsToFile(str string) error {
	file, err := os.Open(keys.filenameoutput)
	if err != nil {
		log.Println("Ошибка записи в файл : " + err.Error())
		return err
	}
	defer file.Close()
	_, er := file.WriteString(str)
	if er != nil {
		log.Fatal(er.Error())
		return err
	}
	return nil
}

func (keys *KeyOfProgramm) DeleteDublicate(rows []string) []string {
	var m map[string]bool = make(map[string]bool)
	for _, v := range rows {
		m[v] = true
	}
	var ans []string
	for k := range m {
		ans = append(ans, k)
	}
	return ans
}

func (keys *KeyOfProgramm) DeleteSpace(rows []string) []string {
	var ans []string
	for _, v := range rows {
		ans = append(ans, strings.TrimSpace(v))
	}
	return ans
}

func (keys *KeyOfProgramm) NumberFirstLessThanNumberSecond(s1, s2 string) bool {
	i, _ := strconv.Atoi(s1)
	j, _ := strconv.Atoi(s2)
	return i < j
}

func (keys *KeyOfProgramm) GetColumnFromRow(row string) (string, error) {
	var str []string = strings.Split(row, " ")
	if keys.column < len(str) {
		return str[keys.column], nil
	}
	return "", errors.New("неправильный индекс")
}

func (keys *KeyOfProgramm) FirstMonthLessThanSecondMonth(m1, m2 string) bool {
	for _, m := range keys.monthset {
		if m == m1 {
			return true
		}
		if m == m2 {
			return false
		}
	}
	return false
}

func (keys *KeyOfProgramm) SortedRows(rows []string) (string, error) {
	var tmperr error = errors.New("вспомогательная ошибка")
	var iserror *error = &tmperr
	if keys.ignorespace {
		rows = (*keys).DeleteSpace(rows)
	}
	if keys.unique {
		rows = (*keys).DeleteDublicate(rows)
	}
	var ans []string
	switch {
	case keys.column > 0:
		switch {
		case keys.month:
			sort.SliceStable(rows, func(i, j int) bool {
				s1, e1 := keys.GetColumnFromRow(rows[i])
				if e1 != nil {
					if iserror != nil {
						iserror = &e1
					}
					log.Fatal("Ошибка" + e1.Error() + " получения столбца " + strconv.Itoa(keys.column))
				}
				s2, e2 := keys.GetColumnFromRow(rows[j])
				if e2 != nil {
					if iserror != nil {
						iserror = &e2
					}
					log.Fatal("Ошибка" + e2.Error() + " получения столбца " + strconv.Itoa(keys.column))
				}
				return keys.FirstMonthLessThanSecondMonth(s1, s2)
			})
		case keys.numberSymbol:
			sort.SliceStable(rows, func(i, j int) bool {
				s1, e1 := keys.GetColumnFromRow(rows[i])
				if e1 != nil {
					if iserror != nil {
						iserror = &e1
					}
					log.Fatal("Ошибка" + e1.Error() + " получения столбца " + strconv.Itoa(keys.column))
				}
				s2, e2 := keys.GetColumnFromRow(rows[j])
				if e2 != nil {
					if iserror != nil {
						iserror = &e2
					}
					log.Fatal("Ошибка" + e2.Error() + " получения столбца " + strconv.Itoa(keys.column))
				}
				return keys.NumberFirstLessThanNumberSecond(s1, s2)
			})
		default:
			sort.SliceStable(rows, func(i, j int) bool {
				s1, e1 := keys.GetColumnFromRow(rows[i])
				if e1 != nil {
					if iserror != nil {
						iserror = &e1
					}
					log.Fatal("Ошибка" + e1.Error() + " получения столбца " + strconv.Itoa(keys.column))
				}
				s2, e2 := keys.GetColumnFromRow(rows[j])
				if e2 != nil {
					if iserror != nil {
						iserror = &e2
					}
					log.Fatal("Ошибка" + e2.Error() + " получения столбца " + strconv.Itoa(keys.column))
				}
				return s1 < s2
			})
		}
		ans = rows
	case keys.month:
		sort.SliceStable(rows, func(i, j int) bool {
			return keys.FirstMonthLessThanSecondMonth(rows[i], rows[j])
		})
		ans = rows
	case keys.numberSymbol:
		sort.SliceStable(rows, func(i, j int) bool {
			return keys.NumberFirstLessThanNumberSecond(rows[i], rows[j])
		})
		ans = rows
	case keys.sorted:
		issort := sort.SliceIsSorted(rows, func(i, j int) bool {
			return rows[i] < rows[j]
		})
		if !issort {
			return "файл был отсортирован", nil
		}
		return "файл не был отсортирован", nil
	default:
		sort.SliceStable(rows, func(i, j int) bool {
			return rows[i] < rows[j]
		})
		ans = rows
	}
	if keys.reverse {
		var tmp []string
		for x := 0; x < len(ans); x++ {
			tmp = append(tmp, ans[(len(ans)-1)-x])
		}
		ans = tmp
	}
	var sb strings.Builder
	for _, s := range ans {
		sb.WriteString(s + "\n")
	}
	if (*iserror).Error() != tmperr.Error() {
		return sb.String(), *iserror
	}
	return sb.String(), nil
}

func main() {
	var keys KeyOfProgramm
	keys.KeyOfProgramm()
	keys.monthset = []string{"Jan", "Feb", "Mar", "Apr", "May",
		"June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}
	s, err := keys.ReadRowsFromFile()
	if err != nil {
		log.Fatal(err)
	}
	r, e := keys.SortedRows(s)
	if e != nil {
		log.Fatal(e)
	}
	if keys.output {
		er := keys.SaveRowsToFile(r)
		if er != nil {
			log.Fatal(er)
		}
	} else {
		fmt.Println(r)
	}
}
