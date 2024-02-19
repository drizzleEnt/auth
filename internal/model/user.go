package model

type UserInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
