package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, filename string) Store {
	switch store {
	case FileType:
		return &fileStore{filename}
	default:
		return nil
	}
}

type fileStore struct {
	FilePath string
}

func (f *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.FilePath, fileData, 0644)
}

func (f *fileStore) Read(data interface{}) error {
	fileData, err := os.ReadFile(f.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, &data)
}
