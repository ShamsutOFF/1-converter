package main

import (
	"1-converter/3-struct/api"
	"1-converter/3-struct/bins"
	"1-converter/3-struct/config"
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
	fileStorage := storage.NewFileStorage("storage")

	// Сохраняем данные
	binList.AddBin(*bins.NewBin("NewBin", "content", false))
	binList.AddBin(*bins.NewBin("NewBin2", "content2", false))

	// Загружаем конфиг из .env (нужно будет KEY=... в .env)
	cfg := config.NewConfig()

	// Передаём конфиг в API
	myApi := api.NewApi(cfg)
	fmt.Println("API initialized with key:", myApi.Config.Key)

	err := fileStorage.SaveBins(*binList, "data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Читаем данные
	loadedList, err := fileStorage.ReadBins("data.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d bins\n", len(loadedList.Bins))
	fmt.Println(loadedList.Bins)
}
