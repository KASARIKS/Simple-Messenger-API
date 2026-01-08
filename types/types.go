package types

type UserStore interface {
	CreateUser(User) error
	GetUserById(id int) (*User, error)
	DeleteUserById(id int) error
}

type User struct {
	Id       int    `json:"id"`
	Password string `json:"-"`
	Nickname string `json:"nickname"`
}

type RegisterPayload struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type LoginPayload struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}
