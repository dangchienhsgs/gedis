package server

import (
	"net"
	"bufio"
	"fmt"
	"strconv"
)

type Server struct {
	Port int
}

func (server Server) Run() {
	host := ":" + strconv.Itoa(server.Port)
	ln, _ := net.Listen("tcp", host)
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
	}
}

func NewServer(port int) Server {
	return Server{Port: port}
}
