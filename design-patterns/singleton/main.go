// Reference example: https://refactoring.guru/design-patterns/singleton/go/example
package main

import (
	"fmt"
	"sync"
)

type Server struct{}

var lock = &sync.Mutex{}
var server *Server

var (
	MsgCreatingServer       = "Creating server.."
	MsgServerAlreadyCreated = "Server already created" 
)

func NewServer() *Server {
	if server == nil {
		lock.Lock()
		defer lock.Unlock()
		if server == nil {
			fmt.Println(MsgCreatingServer)
			server = &Server{}
		} else {
			fmt.Println(MsgServerAlreadyCreated)
		}
	} else {
		fmt.Println(MsgServerAlreadyCreated)
	}
	return server
}

func main() {
	for i := 0; i < 10; i++ {
		go NewServer()
	}
	fmt.Scanln()
}
