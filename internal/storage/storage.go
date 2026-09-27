package storage

import (
	"fmt"

	"github.com/DmitryVysochin/storage/internal/file"

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
	file, err := file.NewFile(filename, blob)
	if err == nil {
		s.files[file.UUID] = file
		return file, nil
	}
	return nil, err
}

func (s *Storage) GetByUUID(fileUUID uuid.UUID) (*file.File, error) {
	if file, ok := s.files[fileUUID]; ok {
		return file, nil
	} else {
		return nil, fmt.Errorf("file %v not found", fileUUID)
	}
}
