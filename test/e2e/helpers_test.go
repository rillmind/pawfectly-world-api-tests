package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

type User struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func createUser(t *testing.T) (*User, error) {
	t.Helper()

	var user User

	body := `{
	"nome": "user",
	"username": "user",
	"email": "user@email.com",
	"senha": "123456"
	}`

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Post(baseUrl+"/user", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatalf("Erro ao tentar deletar usuario: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("status-code: %d, want: %d", resp.StatusCode, http.StatusCreated)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return &user, err
	}

	return &user, nil
}

func deleteUser(t *testing.T, id string) {
	t.Helper()

	req, err := http.NewRequest(http.MethodDelete, baseUrl+"/user/"+id, bytes.NewBufferString(""))
	if err != nil {
		t.Fatalf("Erro ao tntar deletar usuario: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BEARER+authToken)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Erro ao executar requisicao DELETE /user/:id: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status-code: %d, want %d", resp.StatusCode, http.StatusNoContent)
	}

	defer resp.Body.Close()
}
