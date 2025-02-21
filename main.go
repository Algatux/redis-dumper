package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

const HASH_KEY = "alga.algatux.dev"

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379", DB: 0})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Errore di connessione a Redis: %v", err)
	}
	fmt.Println("Connesso a Redis:", pong)

	hashMap, err := rdb.HGetAll(ctx, HASH_KEY).Result()
	if err != nil {
		log.Fatalf("Errore nel recupero della hashmap: %v", err)
	}

	for key, value := range hashMap {
		fmt.Printf("HSET %s %s %s\n", HASH_KEY, key, strconv.Quote(value))
	}
}
