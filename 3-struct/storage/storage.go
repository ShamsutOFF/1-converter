package storage

import (
	"1-converter/3-struct/bins"
	"1-converter/3-struct/file"
	"encoding/json"
	"os"
	"path/filepath"
)

type Storage interface {
	SaveBins(list bins.BinList, filename string) error
	ReadBins(filename string) (bins.BinList, error)
}

type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	// создаём папку если её нет
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return &FileStorage{path: path}
}

func (fs *FileStorage) SaveBins(list bins.BinList, filename string) error {
	bytes, err := json.MarshalIndent(list, "", "  ") // красивый JSON
	if err != nil {
		return err
	}

	fullPath := filepath.Join(fs.path, filename)
	return file.WriteFile(bytes, fullPath)
}

func (fs *FileStorage) ReadBins(filename string) (bins.BinList, error) {
	fullPath := filepath.Join(fs.path, filename)

	bytes, err := file.ReadFile(fullPath)
	if err != nil {
		return bins.BinList{}, err
	}
	err = file.CheckIsOurJSON(bytes)
	if err != nil {
		return bins.BinList{}, err
	}
	var list bins.BinList
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return bins.BinList{}, err
	}
	return list, nil
}
