package cache

import (
	"context"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	cacheConfig := &RedisConfig{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	client := New(cacheConfig)
	err := client.cli.Set(context.Background(), "God", "Yao", 0).Err()
	if err != nil {
		t.Error(err.Error())
	}
	if err = client.Set("Mike", "Join", 0); err != nil {
		t.Error(err.Error())
	}
	result, _ := client.cli.Get(context.Background(), "God").Result()
	log.Println(result)
}
