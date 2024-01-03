package main

import (
	"anystorage/storage"
	"fmt"
)

func main() {
	s := storage.New(storage.WithMemoryRepository())
	fmt.Println(s)
}
