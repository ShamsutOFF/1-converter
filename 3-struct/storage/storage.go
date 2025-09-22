package storage

import (
	"1-converter/3-struct/bins"
	"1-converter/3-struct/file"
	"encoding/json"
)

// Storage интерфейс
type Storage interface {
	SaveBins(list bins.BinList, filename string) error
	ReadBins(filename string) (bins.BinList, error)
}

// FileStorage реализация интерфейса для работы с файлами
type FileStorage struct {
	path string
}

// NewFileStorage конструктор
func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
}

func (fs *FileStorage) SaveBins(list bins.BinList, filename string) error {
	bytes, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = file.WriteFile(bytes, filename)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStorage) ReadBins(filename string) (bins.BinList, error) {
	bytes, err := file.ReadFile(filename)
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
