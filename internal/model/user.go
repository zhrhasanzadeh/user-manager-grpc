package model

import "time"

type User struct {
	Id            int       `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Email         string    `json:"email"`
	MobileNo      string    `json:"mobile_no"`
	Birthdate     time.Time `json:"birthdate"`
	RegisterDate  time.Time `json:"register_date"`
	LastLoginTime time.Time `json:"last_login_time"`
}
