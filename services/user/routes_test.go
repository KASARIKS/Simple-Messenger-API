package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kasariks/simple_messenger_api/services/auth"
	"github.com/kasariks/simple_messenger_api/types"
)

const (
	correctPassword = "123"
	existingUserId  = "first"
)

func TestUserServiceRegisterHandler(t *testing.T) {
	userStore := &mockUserStore{}
	userHandler := NewHandler(userStore)

	t.Run("must register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Id:       "test",
			Password: correctPassword,
			Nickname: "test",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/register", userHandler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("must fail to register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Id:       existingUserId,
			Password: correctPassword,
			Nickname: "test",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			log.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/register", userHandler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

func TestUserServiceLoginHandler(t *testing.T) {
	userStore := &mockUserStore{}
	userHandler := NewHandler(userStore)

	t.Run("must login the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Id:       existingUserId,
			Password: correctPassword,
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/login", userHandler.HandleLogin)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected code %d, got %d", rr.Code, http.StatusOK)
		}
	})

	t.Run("must forbid to login", func(t *testing.T) {
		payloads := []types.LoginUserPayload{
			{
				Id:       existingUserId,
				Password: "incorrectPassword",
			},
			{
				Id:       "notExistingUserId",
				Password: correctPassword,
			},
		}

		expectedCodes := []int{
			http.StatusNotAcceptable,
			http.StatusBadRequest,
		}

		for i, payload := range payloads {
			marshalled, _ := json.Marshal(payload)
			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(marshalled))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := http.NewServeMux()

			router.HandleFunc("/login", userHandler.HandleLogin)
			router.ServeHTTP(rr, req)

			if rr.Code != expectedCodes[i] {
				t.Errorf("Expected code %d, got %d", rr.Code, expectedCodes[i])
			}
		}
	})
}

type mockUserStore struct {
}

func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func (s *mockUserStore) GetUserById(id string) (*types.User, error) {
	hashed, err := auth.HashPassword("123")
	if err != nil {
		return nil, err
	}

	if id == "first" {
		return &types.User{
			Id:             id,
			HashedPassword: hashed,
		}, nil
	}

	return nil, errors.New("empty row")
}

func (s *mockUserStore) DeleteUserById(id string) error {
	return nil
}
