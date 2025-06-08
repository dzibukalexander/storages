package storagesssssss

import (
	"fmt"

	"github.com/dzibukalexander/storages/v2/internal/storage"
)

func NewStorage() *storage.Storage {
	fmt.Println("tag 2.1.0")
	return storage.NewStorage()
}
