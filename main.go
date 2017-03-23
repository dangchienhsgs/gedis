package main

import (
	"fmt"
	"gedis/server"
	"strconv"
)

func main() {
	fmt.Println("Start a server at " + strconv.Itoa(3000))
	s := server.Server{Port: 3000}
	s.Run()
}
