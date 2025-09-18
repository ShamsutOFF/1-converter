package file

import (
	"1-converter/3-struct/bins"
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CheckIsOurJSON(file []byte) error {
	err := json.Unmarshal(file, &bins.BinList{})
	if err != nil {
		fmt.Println("JSON parsing error", err)
		return err
	}
	return nil
}

func WriteFile(content []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	fmt.Println("File written successfully")
	return nil
}
