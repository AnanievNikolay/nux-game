package domain

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}
