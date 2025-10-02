package api_test

import (
	"log"
	"os"
	"testing"

	"1-converter/3-struct/api"
	"1-converter/3-struct/bins"
	"1-converter/3-struct/config"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	possiblePaths := []string{".env", "../.env", "../../.env"}
	for _, p := range possiblePaths {
		if err := godotenv.Load(p); err == nil {
			break
		}
	}
	if os.Getenv("KEY") == "" {
		log.Fatal("❌ API key is empty! Проверь .env")
	}

	code := m.Run()
	os.Exit(code)
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
func createTempBin(t *testing.T, a *api.Api) *bins.Bin {
	b := bins.NewBin("test_bin", "hello world", false)

	created, err := a.CreateBin(*b)
	if err != nil {
		t.Fatalf("failed to create bin: %v", err)
	}
	t.Logf("created test bin ID=%s", created.ID)
	return created
}

func deleteTempBin(t *testing.T, a *api.Api, id string) {
	if err := a.DeleteBin(id); err != nil {
		t.Errorf("cleanup delete failed for Bin=%s: %v", id, err)
	}
}

// --- тесты ---

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

	deleteTempBin(t, apiClient, created.ID)
}
