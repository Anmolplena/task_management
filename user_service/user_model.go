package user_service

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Type     string `json:"type"` // default or admin
}