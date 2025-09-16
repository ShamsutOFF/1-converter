package main

import (
	"fmt"
	"time"
)

// Bin представляет один бинк (файл в облаке)
type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Content   string    `json:"content,omitempty"` // Содержимое файла
}

// BinList представляет список бинков
type BinList struct {
	Bins []Bin `json:"bins"`
}

// Генерация уникального ID (простой вариант)
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// AddBin Добавление бинка в список
func (bl *BinList) AddBin(bin Bin) {
	bl.Bins = append(bl.Bins, bin)
}

// FindByID Поиск бинка по ID
func (bl *BinList) FindByID(id string) *Bin {
	for _, bin := range bl.Bins {
		if bin.ID == id {
			return &bin
		}
	}
	return nil
}

// NewBin Конструктор для создания нового Bin
func NewBin(name string, content string, private bool) *Bin {
	return &Bin{
		ID:        generateID(), // нужно реализовать
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
		Content:   content,
	}
}

// NewBinList Конструктор для BinList
func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}

func main() {
	bin := NewBin("asd", "qwe", false)
	fmt.Println(bin)
	binList := NewBinList()
	binList.AddBin(*bin)
	binList.AddBin(*bin)
	fmt.Println(binList)
}
