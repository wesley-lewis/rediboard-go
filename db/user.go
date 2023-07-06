package db

import (
    "fmt"

    "github.com/go-redis/redis/v8"
)

type User struct {
    Username        string      `json:"username" binding:"required"`
    Points          int         `json:"points" binding:"required"`
    Rank            int         `json:"rank"`
}

func (db *Database) SaveUser(user *User) error {

}
