package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	m := []string{"Jan", "Feb", "Mar", "Apr", "May",
		"June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}

	unitTests := []struct {
		name      string
		keys      KeyOfProgramm
		output    string
		iserror   bool
		typeerror string
	}{
		{"sort test1", KeyOfProgramm{filenameinput: "test1"}, "1\n2\n3\n4\n4\n6\n", false, ""},
		{"sort -r test1", KeyOfProgramm{filenameinput: "test1", reverse: true}, "6\n4\n4\n3\n2\n1\n", false, ""},
		{"sort -k1 test2", KeyOfProgramm{filenameinput: "test2", column: 1}, "1 2\n3 4\n5 6\n7 8\n", false, ""},
		{"sort -r -k1 test2", KeyOfProgramm{filenameinput: "test2", reverse: true, column: 1}, "7 8\n5 6\n3 4\n1 2\n", false, ""},
		{"sort -k2 test2", KeyOfProgramm{filenameinput: "test2", column: 2}, "", true, "неправильный индекс"},
		{"sort -r -k2 test2", KeyOfProgramm{filenameinput: "test2", reverse: true, column: 2}, "", true, "неправильный индекс"},
		{"sort -u test3", KeyOfProgramm{filenameinput: "test3", unique: true}, "1 4\n2 5\n5 6\n", false, ""},
		{"sort -n test4", KeyOfProgramm{filenameinput: "test4", numberSymbol: true}, "-2\n1\n3\n", false, ""},
		{"sort -b test4", KeyOfProgramm{filenameinput: "test5", ignorespace: false}, " 1 \n1\n2\n", false, ""},
		{"sort -M test4", KeyOfProgramm{filenameinput: "test6", month: true, monthset: m}, "Jan\nAug\nDec\n", false, ""},
	}

	for _, test := range unitTests {
		t.Run(test.name, func(t *testing.T) {
			rows, _ := test.keys.ReadRowsFromFile()
			ans, err := test.keys.SortedRows(rows)
			if test.iserror {
				if err != nil {
					if err.Error() != test.typeerror {
						t.Errorf("была выброшена не та ошибка, требуемая ошибка %v имеющаяся ошибка %v", test.typeerror, err.Error())
					}
				} else {
					t.Errorf("должна была быть выброшена ошибка %v, но приложение не выбросило ошибок", test.typeerror)
				}
			} else {
				if err != nil {
					t.Errorf("была выброшена ошибка %v", err.Error())
				}
				if ans != test.output {
					t.Errorf("результат действия программы (%v) не удовлетворяет выходными данными (%v)", ans, test.output)
				}
			}
		})
	}
}
