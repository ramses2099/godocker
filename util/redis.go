package util

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

func InitializeRedis() {
	client = redis.NewClient((&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}))

}

func SetBookList(bookList *Collection) {
	err := client.Set(ctx, "booklist", bookList, 0).Err()
	if err != nil {
		log.Println("Couldn't save the bookList, err ", err)
	}
}

func GetBookList() *Collection {
	val, err := client.Get(ctx, "booklist").Result()
	if err != nil {
		log.Println("No have boooklist ", err)
	}

	var books Collection
	err = json.Unmarshal([]byte(val), &books)
	if err != nil {
		panic(err)
	}
	return &books
}
