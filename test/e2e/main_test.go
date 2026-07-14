package e2e

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

const (
	BEARER string = "Bearer "
)

var (
	baseUrl   string
	client    *http.Client
	authToken string
)

func TestMain(m *testing.M) {
	err := godotenv.Load()

	baseUrl = os.Getenv("API_BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://127.0.0.1:2306"
	}

	token, err := loginNGetToken()
	if err != nil {
		panic("Falha ao autenticar na api: " + err.Error())
	}

	authToken = token

	client = &http.Client{
		Timeout: time.Second * 10,
	}

	code := m.Run()
	os.Exit(code)
}

func loginNGetToken() (string, error) {
	body := `{
	"email_ou_username": "adm",
	"senha": "123456"
	}`
	resp, err := http.Post(baseUrl+"/auth", "application/json", bytes.NewBufferString(body))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	token := resp.Header.Get("Authorization")

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Status: %d, want %d", resp.StatusCode, http.StatusOK)
	}

	return token, nil
}
