package api

import (
	"1-converter/3-struct/bins"
	"1-converter/3-struct/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Api struct {
	Config *config.Config
	Client *http.Client
}

func NewApi(cfg *config.Config) *Api {
	return &Api{
		Config: cfg,
		Client: &http.Client{Timeout: 15 * time.Second}, // ⬅️ теперь максимум 15 сек
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
// GetBin получает Bin с retry и расширенными логами
func (a *Api) GetBin(id string) (*bins.Bin, error) {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)

	var lastErr error
	for attempt := 1; attempt <= 3; attempt++ {
		fmt.Printf("➡️ [GetBin] Attempt %d: URL=%s\n", attempt, url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("request build error: %v", err)
		}
		req.Header.Set("X-Master-Key", a.Config.Key)

		resp, err := a.Client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("request error: %v", err)
			fmt.Println("⚠️ [GetBin] request failed:", lastErr)
			time.Sleep(time.Second) // пауза перед повтором
			continue
		}

		body, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close() // закрываем сразу
		fmt.Printf("➡️ [GetBin] Status=%d; len(body)=%d\n", resp.StatusCode, len(body))

		if err != nil {
			lastErr = fmt.Errorf("read body error: %v", err)
			fmt.Println("⚠️ [GetBin] body read failed:", lastErr)
			time.Sleep(time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
			fmt.Println("⚠️ [GetBin] bad status:", lastErr)
			time.Sleep(time.Second)
			continue
		}

		var result struct {
			Record bins.Bin `json:"record"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			lastErr = fmt.Errorf("unmarshal error: %v, raw=%s", err, string(body))
			fmt.Println("⚠️ [GetBin] JSON parse failed:", lastErr)
			time.Sleep(time.Second)
			continue
		}

		result.Record.ID = id
		fmt.Println("✅ [GetBin] Success:", result.Record)
		return &result.Record, nil
	}

	return nil, fmt.Errorf("GetBin failed after 3 attempts: %v", lastErr)
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
