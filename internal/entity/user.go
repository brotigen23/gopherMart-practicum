package entity

type User struct {
	ID       int     `json:"id"`
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
