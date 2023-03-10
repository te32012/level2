package main

import "testing"

func TestGrep(t *testing.T) {
	testtable := []struct {
		testconf     ConfigGrep
		testname     string
		out          string
		iserror      bool
		messageerror string
	}{
		{testconf: ConfigGrep{inputfile: "test1", pattern: "name"}, testname: "grep name test1", out: "my name is noname\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test1", pattern: "name", invert: true}, testname: "grep -v name test1", out: "li st symbols\nmy work is programmer\ni am programmer\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test1", pattern: "name", invert: true, count: true}, testname: "grep -v -c name test1", out: "3\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test1", line: 3}, testname: "grep -n3 name test1", out: "i am programmer", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test1", line: 4}, testname: "grep -n4 name test1", out: "", iserror: true, messageerror: "количество строк в файле меньше индекса указанной строки"},
		{testconf: ConfigGrep{inputfile: "test1", pattern: "name", before: 2}, testname: "grep -B2 name test1", out: "li st symbols\nmy name is noname\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test1", pattern: "name", after: 2}, testname: "grep -A2 name test1", out: "my name is noname\nmy work is programmer\ni am programmer\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test1", pattern: "name", context: 1}, testname: "grep -C1 name test1", out: "li st symbols\nmy name is noname\nmy work is programmer\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test2", pattern: "adobe", ignore: true}, testname: "grep -i name test2", out: "adobe\naDoBe\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test2", pattern: "adobe", ignore: true}, testname: "grep -i name test2", out: "adobe\naDoBe\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test3", pattern: "^[a-z]*"}, testname: "grep '^[a-z]*' test3", out: "^[a-z]*\nabc\nz\n", iserror: false, messageerror: ""},
		{testconf: ConfigGrep{inputfile: "test3", pattern: "^[a-z]*", fixed: true}, testname: "grep -F '^[a-z]*' test3", out: "^[a-z]*\n", iserror: false, messageerror: ""},
	}
	for _, testcase := range testtable {
		rows, er := testcase.testconf.ReadRowsFromFile()
		if er != nil {
			t.Errorf("Ошибка чтения файла %v", er.Error())
		}
		ans, err := testcase.testconf.Grep(rows)
		if testcase.iserror {
			if err != nil {
				if err.Error() != testcase.messageerror {
					t.Errorf("Должна была быть ошибка %v но выброшенна ошибка %v", testcase.messageerror, err.Error())
				}
			} else {
				t.Errorf("Должна была быть выброшенна ошибка %v но никакой ошибки выброшенно не было", testcase.iserror)
			}
		} else {
			if err != nil {
				t.Errorf("Была выброшена ошибка %v но ошибок быть не должно", err.Error())
			}
			if testcase.out != ans {
				t.Errorf("Был получен результат %v требуется результат %v", ans, testcase.out)
			}
		}
	}
}
