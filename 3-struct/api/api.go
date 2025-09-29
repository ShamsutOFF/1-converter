package api

import (
	"1-converter/3-struct/bins"
	"1-converter/3-struct/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	Config *config.Config
	Client *http.Client
}

func NewApi(cfg *config.Config) *Api {
	return &Api{
		Config: cfg,
		Client: &http.Client{},
	}
}

// CreateBin Создание Bin
func (a *Api) CreateBin(bin bins.Bin) (*bins.Bin, error) {
	url := "https://api.jsonbin.io/v3/b"

	data, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	fmt.Println("KEY = ", a.Config.Key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.Config.Key)

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var result struct {
		Record   bins.Bin `json:"record"`
		Metadata struct {
			ID string `json:"id"`
		} `json:"metadata"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// Подставляем ID
	result.Record.ID = result.Metadata.ID

	return &result.Record, nil
}

// GetBin Получение Bin
func (a *Api) GetBin(id string) (*bins.Bin, error) {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Master-Key", a.Config.Key)

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var result struct {
		Record bins.Bin `json:"record"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	result.Record.ID = id
	return &result.Record, nil
}

// UpdateBin Обновление Bin
func (a *Api) UpdateBin(id string, bin bins.Bin) error {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)

	data, err := json.Marshal(bin)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.Config.Key)

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(body))
	}
	return nil
}

// DeleteBin Удаление Bin
func (a *Api) DeleteBin(id string) error {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Master-Key", a.Config.Key)

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(body))
	}
	return nil
}
