package main

import (
	"testing"
)

func TestUnpacStr(t *testing.T) {
	tests := []struct {
		nameoftests string
		input       string
		output      string
		iserror     bool
	}{
		{nameoftests: "string1", input: "b2c4e", output: "bbcccce", iserror: false},
		{nameoftests: "string2", input: "abcd", output: "abcd", iserror: false},
		{nameoftests: "error1", input: "443", output: "", iserror: true},
		{nameoftests: "empty", input: "", output: "", iserror: false},
		{nameoftests: "escape", input: "qwerty\\1\\5", output: "qwerty15", iserror: false},
		{nameoftests: "escape2", input: "qwerty\\15", output: "qwerty11111", iserror: false},
		{nameoftests: "escape3", input: "qwerty\\\\3", output: "qwerty\\\\\\", iserror: false},
	}
	for _, struc := range tests {
		t.Run(struc.nameoftests, func(t *testing.T) {
			res, err := UnpacStr(struc.input)
			if struc.iserror {
				if err != nil {
					if err.Error() != "uncorect string" {
						t.Errorf("должна быть ошибка - строка не корректна, но есть ошибка %v", err.Error())
					}
				} else {
					t.Errorf("должна быть ошибка - строка не корректна, но ошибки нет")
				}
			} else {
				if err != nil {
					t.Errorf("не должно быть ошибки, но есть ошибка %v", err.Error())
				}
				if res != struc.output {
					t.Errorf("нужный результат %v, имеющийся результат %v", struc.output, res)
				}
			}
		})
	}
}
