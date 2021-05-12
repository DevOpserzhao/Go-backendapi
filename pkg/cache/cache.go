package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"log"
	"time"
)

type RedisFace interface {
	Set(string, string, time.Duration) error
	Get(string) (string, error)
	SAdd(string, ...interface{}) error
	SIsMember(string, interface{}) bool
	SClear(string) bool
}

func NewRedisFace(client *RedisClient) RedisFace {
	var rf RedisFace = client
	return rf
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type RedisClient struct {
	cli *redis.Client
}

func New(config *RedisConfig) *RedisClient {
	return &RedisClient{cli: open(config)}
}

func open(c *RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Println("\033[1;31;31m=========== Redis ============\033[0m")
		log.Println("\033[1;34;34m	  Please Check Redis\033[0m")
		log.Println("\033[1;31;31m==============================\033[0m")
		return nil
	}
	log.Println()
	log.Printf("\033[1;32;32m Redis RUNING [%s] \033[0m", c.Addr)
	log.Println()
	return client
}

func (client *RedisClient) Set(key string, value string, expiration time.Duration) error {
	return client.cli.Set(context.Background(), key, value, expiration).Err()
}

var errorRedisKeyNotExist = errors.New("Redis Key Not Exist")

const emptyString = ""

func (client *RedisClient) Get(key string) (string, error) {
	value, err := client.cli.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return emptyString, errorRedisKeyNotExist
	}
	if err != nil && err != redis.Nil {
		return emptyString, err
	}
	return value, nil
}

func (client *RedisClient) SAdd(key string, members ...interface{}) error {
	return client.cli.SAdd(context.Background(), key, members...).Err()
}

func (client *RedisClient) SIsMember(key string, member interface{}) bool {
	return client.cli.SIsMember(context.Background(), key, member).Val()
}

func (client *RedisClient) SClear(key string) bool {
	return client.cli.Expire(context.Background(), key, 0).Val()
}
