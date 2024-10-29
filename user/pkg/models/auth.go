package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	UserName string `json:"user_name" go`
	Password string `json:"password"`
}
