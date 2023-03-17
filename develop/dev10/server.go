package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type ConfigTelnetServer struct {
	Host string
	Port int
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
		conn.SetDeadline(time.Now().Add(time.Millisecond * time.Duration(700)))
		if err != nil {
			log.Fatal("ошибочное подключение к серверу")
		}
		defer conn.Close()
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	for {
		tmp := make([]byte, 2^30)
		hello := "привет я сервер"
		conn.Write([]byte(hello))
		conn.Read(tmp)
		os.Stdout.Write(tmp)
	}
}
func main() {
	var server ConfigTelnetServer = ConfigTelnetServer{}
	server.ConfigTelnetServer("", 2023)
	server.StartServer()
}
