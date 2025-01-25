package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignup struct {
	UserLogin
	UserType int `json:"user_type"`
}
