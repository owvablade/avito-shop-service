package model

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Coins    int    `db:"coins"`
}

type MerchItem struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Price int    `db:"price"`
}

type UserInfo struct {
	Coins       int         `json:"coins"`
	Inventory   []Inventory `json:"inventory"`
	CoinHistory CoinHistory `json:"coinHistory"`
}

type Inventory struct {
	Type     string `json:"type" db:"type"`
	Quantity int    `json:"quantity" db:"quantity"`
}

type CoinHistory struct {
	Received []ReceivedTransaction `json:"received"`
	Sent     []SentTransaction     `json:"sent"`
}

type ReceivedTransaction struct {
	FromUser string `json:"fromUser" db:"username"`
	Amount   int    `json:"amount" db:"amount"`
}

type SentTransaction struct {
	ToUser string `json:"toUser" db:"username"`
	Amount int    `json:"amount" db:"amount"`
}
