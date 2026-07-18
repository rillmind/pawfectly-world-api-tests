package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	user := &User{}
	t.Cleanup(func() { deleteUser(t, user.ID) })

	body := `{
	"nome": "user",
	"username": "user",
	"email": "user@email.com",
	"senha": "123456"
	}`

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Post(baseUrl+"/user", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatalf("Erro executar tentar criar usuario: %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatalf("Erro ao decodificar corpo da resposta: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("status-code %d, want: %d", resp.StatusCode, http.StatusCreated)
	}
}

func TestListUsers(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, baseUrl+"/user", bytes.NewBufferString(""))
	if err != nil {
		t.Errorf("Erro ao listar usuarios: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BEARER+authToken)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Erro ao executar requisicao: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status-code: %d, want: %d", resp.StatusCode, http.StatusOK)
	}
}

func TestListUserByID(t *testing.T) {
	user, err := createUser(t)
	if err != nil {
		t.Fatalf("Erro ao crir usuario: %v", err)
	}

	t.Cleanup(func() { deleteUser(t, user.ID) })

	req, err := http.NewRequest(http.MethodGet, baseUrl+"/user/"+user.ID, bytes.NewBufferString(""))
	if err != nil {
		t.Errorf("Erro ao listar usuario por ID: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BEARER+authToken)

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Erro ao tentar executar requisicao: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status-code: %d, want: %d", resp.StatusCode, http.StatusOK)
	}
}

func TestUpdateUserByID(t *testing.T) {
	user, err := createUser(t)
	if err != nil {
		t.Fatalf("Erro ao criar usuario: %v", err)
	}

	t.Cleanup(func() { deleteUser(t, user.ID) })

	body := `{
	"nome": "userUpdated",
	"username": "userUpdated",
	"email": "userUpdated@email.com",
	"senha": "654321"
	}`

	req, err := http.NewRequest(http.MethodPatch, baseUrl+"/user/"+user.ID, bytes.NewBufferString(body))
	if err != nil {
		t.Errorf("Erro ao tentar atualizar usuario com ID. err: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BEARER+authToken)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Erro ao executar requisicao PATCH /user/:id. Err: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("status-code: %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
}

func TestDeleteUserByID(t *testing.T) {
	user, err := createUser(t)
	if err != nil {
		t.Fatalf("Erro ao criar usuario: %v", err)
	}

	req, err := http.NewRequest(http.MethodDelete, baseUrl+"/user/"+user.ID, bytes.NewBufferString(""))
	if err != nil {
		t.Fatalf("Erro ao tentar deletar usuario: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", BEARER+authToken)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Erro ao executar requisicao DELETE /user/:id: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("status-code: %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
}
