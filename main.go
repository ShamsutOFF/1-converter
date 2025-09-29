package main

import (
	"1-converter/3-struct/api"
	"1-converter/3-struct/bins"
	"1-converter/3-struct/config"
	"1-converter/3-struct/storage"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Читаем флаги
	action := flag.String("action", "", "Action: create|get|update|delete|list")
	id := flag.String("id", "", "Bin ID")
	fileName := flag.String("file", "", "Path to json file for create/update")
	flag.Parse()
	fmt.Println("Action:", *action)
	fmt.Println("ID:", *id)
	fmt.Println("File:", *fileName)

	// 2. Создаём конфиг и API клиент
	cfg := config.NewConfig()
	apiClient := api.NewApi(cfg)

	// 3. Локальное хранилище (для list/delete)
	fileStorage := storage.NewFileStorage("file_storage")

	// 4. Определяем действие
	switch *action {
	case "create":
		if *fileName == "" {
			log.Fatal("❌ file is required for create")
		}

		// читаем JSON из файла для API
		content, err := os.ReadFile(*fileName)
		if err != nil {
			log.Fatal("❌ error reading file:", err)
		}

		var newBin bins.Bin
		err = json.Unmarshal(content, &newBin)
		if err != nil {
			log.Fatal("❌ invalid JSON in file:", err)
		}

		// отправляем на API
		createdBin, err := apiClient.CreateBin(newBin)
		if err != nil {
			log.Fatal("❌ API error:", err)
		}

		fmt.Println("✅ Bin created on API:", createdBin.ID, createdBin.Name)

		// сохраняем только в общий локальный файл (data.json)
		list, _ := fileStorage.ReadBins("data.json")
		list.AddBin(*createdBin)

		err = fileStorage.SaveBins(list, "data.json")
		if err != nil {
			log.Fatal("❌ error saving local storage:", err)
		}

	case "get":
		if *id == "" {
			log.Fatal("❌ id is required for get")
		}

		bin, err := apiClient.GetBin(*id)
		if err != nil {
			log.Fatal("❌ API error:", err)
		}

		fmt.Printf("✅ Bin %s (%s): %s\n", bin.ID, bin.Name, bin.Content)

	case "update":
		if *id == "" || *fileName == "" {
			log.Fatal("❌ id and file are required for update")
		}

		content, err := os.ReadFile(*fileName)
		if err != nil {
			log.Fatal("❌ error reading file:", err)
		}

		var updatedBin bins.Bin
		err = json.Unmarshal(content, &updatedBin)
		if err != nil {
			log.Fatal("❌ invalid JSON:", err)
		}

		err = apiClient.UpdateBin(*id, updatedBin)
		if err != nil {
			log.Fatal("❌ API error:", err)
		}

		fmt.Println("✅ Bin updated:", *id)

	case "delete":
		if *id == "" {
			log.Fatal("❌ id is required for delete")
		}

		err := apiClient.DeleteBin(*id)
		if err != nil {
			log.Fatal("❌ API error:", err)
		}

		// обновляем локальное хранилище — удаляем бин по ID
		list, err := fileStorage.ReadBins("data.json")
		if err == nil {
			newList := bins.NewBinList()
			for _, b := range list.Bins {
				if b.ID != *id {
					newList.AddBin(b)
				}
			}
			_ = fileStorage.SaveBins(*newList, "data.json")
		}

		fmt.Println("✅ Bin deleted:", *id)

	case "list":
		list, err := fileStorage.ReadBins("data.json")
		if err != nil {
			log.Fatal("❌ error reading local storage:", err)
		}
		for _, v := range list.Bins {
			fmt.Printf("ID: %s, Name: %s\n", v.ID, v.Name)
		}

	default:
		fmt.Println("Usage examples:")
		fmt.Println("  go run main.go -action=create -file=bin.json")
		fmt.Println("  go run main.go -action=get -id=12345")
		fmt.Println("  go run main.go -action=update -id=12345 -file=bin.json")
		fmt.Println("  go run main.go -action=delete -id=12345")
		fmt.Println("  go run main.go -action=list")
	}
}
