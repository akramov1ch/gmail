package redis

import (
    "github.com/go-redis/redis/v8"
    "context"
    "log"
)

var RDB *redis.Client

func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    if err := RDB.Ping(context.Background()).Err(); err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }

    log.Println("Connected to Redis")
}
