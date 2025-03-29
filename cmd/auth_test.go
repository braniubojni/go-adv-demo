package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"
	"go/adv-demo/pkg/req"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	TestEmail    = "test123@example.com"
	TestPassword = "Abcd1234$"
	TestName     = "Kikos"
)

func TestRegisterSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    TestEmail,
		Password: TestPassword,
		Name:     TestName,
	})

	res, err := http.Post(ts.URL+"/auth/register", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d got %d", http.StatusOK, res.StatusCode)
	}
	body, err := req.Decode[auth.RegisterResponse](res.Body)
	if err != nil || body.Token == "" {
		t.Fatal("Received invalid token")
	}
}

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    TestEmail,
		Password: TestPassword,
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d got %d", http.StatusOK, res.StatusCode)
	}
	body, err := req.Decode[auth.LoginResponse](res.Body)
	if err != nil || body.Token == "" {
		t.Fatal("Received invalid token")
	}
}

func TestLoginNegative(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "not-existing-email@ok.ru",
		Password: "a b c d 1 234",
	})

	res, _ := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Expected %d got %d", http.StatusUnauthorized, res.StatusCode)
	}
}
