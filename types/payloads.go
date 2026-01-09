package types

import "fmt"

type RegisterUserPayload struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type LoginUserPayload struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func ValidateRegisterUserPayload(payload RegisterUserPayload) error {
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

func ValidateLoginUserPayload(payload LoginUserPayload) error {
	if payload.Id == "" {
		return fmt.Errorf("missing id")
	}
	if payload.Password == "" {
		return fmt.Errorf("missing password")
	}

	return nil
}
