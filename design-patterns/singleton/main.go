// Reference example: https://refactoring.guru/design-patterns/singleton/go/example
package main

import (
	"fmt"
	"sync"
)

type Server struct{}

var lock = &sync.Mutex{}
var server *Server

func NewServer() *Server {
	if server == nil {
		lock.Lock()
		defer lock.Unlock()
		if server == nil {
			fmt.Println("creating server..")
			server = &Server{}
		} else {
			fmt.Println("server already created")
		}
	} else {
		fmt.Println("server already created")
	}
	return server
}

func main() {
	for i := 0; i < 10; i++ {
		go NewServer()
	}
	fmt.Scanln()
}
