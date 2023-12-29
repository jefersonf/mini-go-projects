package main

import (
	"anystorage/storage"
	"fmt"
)

func main() {
	s := storage.NewStorage(storage.WithMemoryRepository())
	fmt.Println(s)
}
