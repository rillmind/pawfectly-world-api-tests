package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var user struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func TestCreateUser(t *testing.T) {
	body := `{
	"nome": "user",
	"username": "user",
	"email": "user@email.com",
	"senha": "123456",
	"role": "user"
	}`

	req, err := http.NewRequest(http.MethodPost, baseUrl+"/user", bytes.NewBufferString(body))
	if err != nil {
		t.Fatalf("Erro ao tentar criar usuario: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BEARER+authToken)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Erro ao executar requisiçao: %v", err)

	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatalf("Erro ao decodificar corpo da resposta: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("status: %d, want %d", resp.StatusCode, http.StatusCreated)
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, baseUrl+"/user/"+user.ID, bytes.NewBufferString(""))
	if err != nil {
		t.Fatalf("Erro ao tentar deletar usuario: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BEARER+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Erro ao executar requisiçao: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("status: %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
}
