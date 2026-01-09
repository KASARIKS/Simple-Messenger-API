package types

type User struct {
	Id             string `json:"id"`
	HashedPassword string `json:"-"`
	Nickname       string `json:"nickname"`
}

type Message struct {
	Id          int    `json:"id"`
	AuthorId    string `json:"authorId"`
	RecipientId string `json:"recipientId"`
	Value       string `json:"value"`
	CreatedAt   string `json:"createdAt"`
}
