package main

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123


Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type ConfigTelnetServer struct {
	Host string
	Port int
}

type ConfigTelnetClient struct {
	Timeout int
	Host    string
	Port    int
}

func (configtelnetserver *ConfigTelnetServer) ConfigTelnetServer(host string, port int) {
	configtelnetserver.Host = host
	configtelnetserver.Port = port
}
func (configtelnetserver *ConfigTelnetServer) StartServer() {
	listener, err := net.Listen("tcp", configtelnetserver.Host+":"+strconv.Itoa(configtelnetserver.Port))
	if err != nil {
		log.Fatal("ошибка соедиения с сервером")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("ошибочное подключение к серверу")
		}
		defer conn.Close()
		hello := "привет я сервер"
		conn.Write([]byte(hello))
		scaner := bufio.NewScanner(conn)
		str := scaner.Text()
		os.Stderr.Write([]byte(str))
	}
}

func (configtelnetclient *ConfigTelnetClient) ConfigTelnetClient() {
	flag.IntVar(&configtelnetclient.Timeout, "timeout", 10, "промежуток через который сбрасывается соединение")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("переданно неправильное количество аргументов")
	}
	configtelnetclient.Host = args[0]
	tmp, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("аргумент порт должен быть числом")
	}
	configtelnetclient.Port = tmp
}

// Вывести в консоль
func (configtelnetclient *ConfigTelnetClient) Write(conection net.Conn, cancel context.CancelFunc) {
	for {
		buffer := make([]byte, 2^20)
		_, err := conection.Read(buffer)
		if err != nil {
			log.Println("не можем прочитать символы с сервера")
			cancel()
			break
		}
		_, err = os.Stdout.Write(buffer)
		if err != nil {
			log.Println("не можем вывести данные на консоль")
			cancel()
			break
		}
	}
}

// Прочитать с консоли
func (configtelnetclient *ConfigTelnetClient) Read(conection net.Conn, cancel context.CancelFunc) {
	conection.Write([]byte("привет я клиент"))
	scan := bufio.NewScanner(os.Stdin)
	for {
		if !scan.Scan() {
			log.Println("не можем прочитать символы из стандартного ввода")
			cancel()
			break
		}
		tmp := scan.Text()
		_, err := conection.Write([]byte(tmp))
		if err != nil {
			log.Println("не можем записать символы на сервер")
			cancel()
			break
		}
	}
}

func main() {
	ctx, cns := context.WithCancel(context.Background())

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGQUIT)
	go func() {
		<-sigch
		cns()
	}()
	var client ConfigTelnetClient = ConfigTelnetClient{}
	client.ConfigTelnetClient()
	adress := client.Host + ":" + strconv.Itoa(client.Port)
	connection, err := net.Dial("tcp", adress)
	if err != nil {
		fmt.Errorf(err.Error())
		log.Fatal("ошибка соединения c сервером")
	}
	defer connection.Close()
	connection.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(client.Timeout)))

	go client.Write(connection, cns)
	go client.Read(connection, cns)
	<-ctx.Done()
	log.Println("клиент закрыт")
}
