package storage

import (
	"fmt"

	"github.com/dzibukalexander/storages/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	file, ok := s.files[fileId]
	if !ok {
		//return nil, &NoSuchFile{}
		// return nil, errors.New(fmt.Sprintf("Error: no such file exists with id %v", fileId))
		return nil, fmt.Errorf("Error: no such file exists with id %v", fileId)
	}

	return file, nil
}
