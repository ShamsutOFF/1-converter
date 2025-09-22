package main

import (
	"1-converter/3-struct/bins"
	"1-converter/3-struct/storage"
	"fmt"
	"log"
)

func main() {
	bin := bins.NewBin("asd", "qwe", false)
	fmt.Println(bin)
	binList := bins.NewBinList()
	binList.AddBin(*bin)

	// Создаем экземпляр хранилища
	fileStorage := storage.NewFileStorage("file_storage")

	// Используем через интерфейс
	var st storage.Storage = fileStorage

	// Сохраняем данные
	binList.AddBin(*bins.NewBin("NewBin", "content", false))

	err := st.SaveBins(*binList, "data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Читаем данные
	loadedList, err := st.ReadBins("data.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d bins\n", len(loadedList.Bins))
	fmt.Println(loadedList.Bins)
}
