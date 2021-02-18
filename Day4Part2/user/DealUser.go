package user

import (
	"encoding/hex"
)

type User struct {
	Email    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var Index int

var UserDB = []User{}

var APISECRET = hex.EncodeToString([]byte("une-cle-de-32-bytes-de-long-là!"))
