package message

import (
	"database/sql"
	"time"

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

func (s *Store) CreateMessage(message types.Message) error {
	_, err := s.db.Exec("INSERT INTO messages (authorId, recipientId, value, createdAt)"+
		"VALUES (:authorId, :recipientId, :value, :createdAt);",
		sql.Named("authorId", message.AuthorId),
		sql.Named("recipientId", message.RecipientId),
		sql.Named("value", message.Value),
		sql.Named("createdAt", time.Now().String()))

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetMessageById(id int) (*types.Message, error) {
	msg := new(types.Message)

	row := s.db.QueryRow("SELECT * FROM messages WHERE id = :id;",
		sql.Named("id", id))

	err := row.Scan(
		&msg.Id,
		&msg.AuthorId,
		&msg.RecipientId,
		&msg.Value,
		&msg.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
func (s *Store) GetMessagesByAuthorId(id int) ([]types.Message, error) {
	rows, err := s.db.Query("SELECT * FROM messages WHERE authorId = :authorId;",
		sql.Named("authorId", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	msgs, err := getMessagesFromRows(rows)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (s *Store) GetMessagesByRecipientId(id int) ([]types.Message, error) {
	rows, err := s.db.Query("SELECT * FROM messages WHERE recipientId = :recipientId;",
		sql.Named("recipientId", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	msgs, err := getMessagesFromRows(rows)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func getMessagesFromRows(rows *sql.Rows) ([]types.Message, error) {
	msgs := []types.Message{}
	var tmpMsg types.Message

	for rows.Next() {
		err := rows.Scan(
			&tmpMsg.Id,
			&tmpMsg.AuthorId,
			&tmpMsg.RecipientId,
			&tmpMsg.Value,
			&tmpMsg.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, tmpMsg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}
