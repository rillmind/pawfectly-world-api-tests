package e2e

import (
	"bytes"
	"fmt"
	"log"
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
	client    *http.Client
	baseUrl   string
	authToken string
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")

	baseUrl = os.Getenv("API_BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://127.0.0.1:2306"
	}

	client = &http.Client{
		Timeout: time.Second * 10,
	}

	health := CheckHealth()
	if health != nil {
		log.Fatalf("API offline: %v", health)
	}

	token, err := loginNGetToken()
	if err != nil {
		log.Fatalf("Falha ao autenticar na api: %s", err.Error())
	}

	authToken = token

	code := m.Run()
	os.Exit(code)
}

func loginNGetToken() (string, error) {
	body := fmt.Sprintf(`{
	"email_ou_username": "%s",
	"senha": "%s"
	}`, os.Getenv("ADM_USERNAME"), os.Getenv("ADM_PASSWORD"))

	if os.Getenv("ADM_USERNAME") == "" || os.Getenv("ADM_PASSWORD") == "" {
		return "", fmt.Errorf("Sem credenciais de login no envfile.")
	}

	resp, err := http.Post(baseUrl+"/auth", "application/json", bytes.NewBufferString(body))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	token := resp.Header.Get("Authorization")

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status-code: %d, want %d", resp.StatusCode, http.StatusOK)
	}

	return token, nil
}

func CheckHealth() error {
	resp, err := client.Get(baseUrl + "/health")

	if err != nil {
		return fmt.Errorf("Erro ao executar requisicao: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status-code: %v, want %v", resp.StatusCode, http.StatusOK)
	}

	return nil
}
