package telegrambotapi

// telegram chat
type Chat struct {
	Id        int64  `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}
