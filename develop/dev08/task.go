package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

go vet
ok

golangci-lint run --enable golint task.go
ok
*/

func executeCommand(partofpipiline string, input bytes.Buffer) (bytes.Buffer, error) {
	tmp := strings.Split(partofpipiline, " ")
	cmd := tmp[0]
	args := tmp[1:]
	switch {
	case strings.Contains(cmd, "ls"):
		cmd := exec.Command("ls", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = &input
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		return out, err
	case strings.Contains(cmd, "cd"):
		cmd := exec.Command("cd", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = &input
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		return out, err
	case strings.Contains(cmd, "pwd"):
		cmd := exec.Command("pwd", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = &input
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		return out, err
	case strings.Contains(cmd, "echo"):
		cmd := exec.Command("echo", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = &input
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		return out, err
	case strings.Contains(cmd, "kill"):
		cmd := exec.Command("kill", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = &input
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		return out, err
	case strings.Contains(cmd, "ps"):
		cmd := exec.Command("ps", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = &input
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		return out, err
	case strings.Contains(cmd, "grep"):
		cmd := exec.Command("grep", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = &input
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		return out, err
	case strings.Contains(cmd, "exit") || strings.Contains(cmd, "quite"):
		return bytes.Buffer{}, errors.New("завершение работы")
	default:
		return bytes.Buffer{}, errors.New("команда не поддерживается")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
L1:
	for {
		fmt.Print("user@user$")
		piplineOfCommand, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		piplineOfCommand = strings.TrimSuffix(piplineOfCommand, "\n")
		comands := strings.Split(piplineOfCommand, " | ")
		var tmp bytes.Buffer
		for i, cmd := range comands {
			ans, err := executeCommand(cmd, tmp)
			if err != nil {
				if err.Error() == "завершение работы" {
					goto L2
				}
				fmt.Fprintln(os.Stderr, "ошибка", err.Error())
				goto L1
			}
			if i == len(comands)-1 {
				os.Stdout.Write(ans.Bytes())
			} else {
				tmp = ans
			}
		}
	}
L2:
}
