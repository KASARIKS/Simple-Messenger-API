package types

type UserStore interface {
	CreateUser(User) error
	GetUserById(id string) (*User, error)
	DeleteUserById(id string) error
}

type MessageStore interface {
	CreateMessage(Message) error
	GetMessageById(id int) (*Message, error)
	GetMessagesByAuthorId(id int) ([]Message, error)
	GetMessagesByRecipientId(id int) ([]Message, error)
}
