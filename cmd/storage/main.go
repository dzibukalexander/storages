package main

import (
	"fmt"
	"log"

	"github.com/dzibukalexander/storages/v2/internal/storage"
	"github.com/google/uuid"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)

	file, err := st.Upload("v1", []byte("hello"))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("file", file)

	fmt.Println(st.GetById(file.ID))

	fmt.Println(st.GetById(uuid.Max))
}
