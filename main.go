package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://red-cuk6fsij1k6c73d4j520:6379")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	ctx := context.Background()

	err = client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
}
