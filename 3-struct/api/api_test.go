package api_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"1-converter/3-struct/api"
	"1-converter/3-struct/bins"
	"1-converter/3-struct/config"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	paths := []string{".env", "../.env", "../../.env"}

	var loadErr error
	for _, p := range paths {
		if err := godotenv.Load(p); err == nil {
			loadErr = nil
			break
		} else {
			loadErr = err
		}
	}
	if os.Getenv("KEY") == "" {
		log.Fatalf("❌ API key is empty! Проверь .env. Last error: %v", loadErr)
	}

	// Запуск тестов
	os.Exit(m.Run())
}

// helper: создаём тестовый API‑клиент
func newTestApi() *api.Api {
	cfg := config.NewConfig()
	if cfg.Key == "" {
		log.Fatal("❌ API key is empty! Проверь .env или переменные окружения")
	}
	return api.NewApi(cfg)
}

// helper: создаём временный bin
func createTempBin(t *testing.T, a *api.Api, name, content string) *bins.Bin {
	var created *bins.Bin
	var err error
	for i := 0; i < 3; i++ { // до 3 попыток
		b := bins.NewBin(name, content, false)
		created, err = a.CreateBin(*b)
		if err == nil {
			break
		}
		t.Logf("retry create bin after error: %v", err)
		time.Sleep(time.Second)
	}
	if err != nil {
		t.Fatalf("failed to create bin: %v", err)
	}
	return created
}

// helper: удаляем bin
func deleteTempBin(t *testing.T, a *api.Api, id string) {
	if err := a.DeleteBin(id); err != nil {
		t.Errorf("failed to delete test bin [%s]: %v", id, err)
	} else {
		t.Logf("🗑️ deleted test bin ID=%s", id)
	}
}

// ----------------- TESTS -----------------

func TestCreateBin(t *testing.T) {
	apiClient := newTestApi()
	b := bins.NewBin("create_test", "data", false)

	created, err := apiClient.CreateBin(*b)
	if err != nil {
		t.Fatalf("CreateBin failed: %v", err)
	}
	if created.ID == "" {
		t.Errorf("expected non-empty ID")
	}
	if created.Name != "create_test" {
		t.Errorf("expected Name=create_test, got %s", created.Name)
	}

	// cleanup
	deleteTempBin(t, apiClient, created.ID)
}

func TestGetBin(t *testing.T) {
	apiClient := newTestApi()
	created := createTempBin(t, apiClient, "get_test", "hello")
	fmt.Println("🟢 TestGetBin: created =", created)

	got, err := apiClient.GetBin(created.ID)
	if err != nil {
		t.Fatalf("❌ GetBin failed: %v", err)
	}
	fmt.Println("🟢 TestGetBin: got =", got)

	if got.Content != created.Content {
		t.Errorf("expected Content=%s, got %s", created.Content, got.Content)
	}

	deleteTempBin(t, apiClient, created.ID)
}

func TestUpdateBin(t *testing.T) {
	apiClient := newTestApi()
	created := createTempBin(t, apiClient, "update_test", "old content")

	// изменяем содержимое
	created.Content = "new content"
	err := apiClient.UpdateBin(created.ID, *created)
	if err != nil {
		t.Fatalf("UpdateBin failed: %v", err)
	}

	// проверяем после обновления
	got, err := apiClient.GetBin(created.ID)
	if err != nil {
		t.Fatalf("GetBin after update failed: %v", err)
	}
	if got.Content != "new content" {
		t.Errorf("expected updated content, got %s", got.Content)
	}

	// cleanup
	deleteTempBin(t, apiClient, created.ID)
}

func TestDeleteBin(t *testing.T) {
	apiClient := newTestApi()
	created := createTempBin(t, apiClient, "delete_test", "data to delete")

	// удаляем
	err := apiClient.DeleteBin(created.ID)
	if err != nil {
		t.Fatalf("DeleteBin failed: %v", err)
	}

	// проверяем что bin действительно удалился
	_, err = apiClient.GetBin(created.ID)
	if err == nil {
		t.Errorf("expected error when getting deleted bin, got none")
	}
}
