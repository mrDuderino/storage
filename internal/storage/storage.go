package storage

import "github.com/mrDuderino/storage/internal/file"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(fileName, blob)
	if err != nil {
		return nil, err
	}
	return newFile, nil
}
