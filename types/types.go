package types

import (
	"fmt"
)

type UserStore interface {
	CreateUser(User) error
	GetUserById(id string) (*User, error)
	DeleteUserById(id string) error
}

type User struct {
	Id             string `json:"id"`
	HashedPassword string `json:"-"`
	Nickname       string `json:"nickname"`
}

type RegisterPayload struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type LoginPayload struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func ValidateRegisterPayload(payload RegisterPayload) error {
	if payload.Id == "" {
		return fmt.Errorf("missing id")
	}
	if payload.Password == "" {
		return fmt.Errorf("missing password")
	}
	if payload.Nickname == "" {
		return fmt.Errorf("missing nickname")
	}

	return nil
}

func ValidateLoginPayload(payload LoginPayload) error {
	if payload.Id == "" {
		return fmt.Errorf("missing id")
	}
	if payload.Password == "" {
		return fmt.Errorf("missing password")
	}

	return nil
}
