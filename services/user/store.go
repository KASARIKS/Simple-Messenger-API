package user

import (
	"database/sql"

	"github.com/kasariks/simple_messenger_api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (id, password, nickname) VALUES (:id, :password, :nickname);",
		sql.Named("id", user.Id),
		sql.Named("password", user.HashedPassword),
		sql.Named("nickname", user.Nickname))
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserById(id string) (*types.User, error) {
	user := new(types.User)

	row := s.db.QueryRow("SELECT * FROM users WHERE id = :id;",
		sql.Named("id", id))

	err := row.Scan(
		&user.Id,
		&user.HashedPassword,
		&user.Nickname,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) DeleteUserById(id string) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = :id;",
		sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}
