package main

import (
	"testing"
)

func TestGetSetOfAnnagrams(t *testing.T) {
	testtable := []struct {
		name   string
		input  []string
		output map[string][]string
	}{
		{name: "first test", input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			output: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}}},
		{name: "second test", input: []string{"пять", "ятьп", "это"},
			output: map[string][]string{"пять": {"пять", "ятьп"}}},
	}

	for _, testcase := range testtable {
		t.Run(testcase.name, func(t *testing.T) {
			ans := GetSetOfAnnagrams(testcase.input)
			for k1, v1 := range testcase.output {
				set, ok := ans[k1]
				if !ok {
					t.Errorf("ключевое слово - анаграмма %v не было найдено", k1)
				}
				if len(set) != len(v1) {
					t.Errorf("Должно быть %v но результат %v", v1, set)
				}
				for x := 0; x < len(set); x++ {
					if set[x] != v1[x] {
						t.Errorf("Должно быть %v но результат %v", v1, set)
					}
				}
			}
		})
	}

}
