package db 

import (
    "context"
    "errors"
    "log"

    "github.com/go-redis/redis/v8"
)

type Database struct {
    Client *redis.Client
}

var (
    ErrNil = errors.New("no matching record for the record found in db")
    Ctx = context.TODO()
)

func NewDatabase(address string) ( *Database, error) {
    client := redis.NewClient(&redis.Options{
        Addr: address,
        Password: "",
        DB: 0,
    })

    if err := client.Ping(Ctx).Err(); err != nil {
        return nil, err 
    }
    log.Println("Database is alive")
    return &Database {
        Client: client,
    }, nil 
}
