package main

import (
	"anystorage/storage"
	"fmt"
)

func main() {
	// creates a new in-memory storage.
	s := storage.New(storage.WithMemoryRepository())
	fmt.Println(s)
}
