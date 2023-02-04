package main

import (
	"fmt"
	"log"

	"github.com/mrDuderino/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello there"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfuly uploaded:", file)
}
