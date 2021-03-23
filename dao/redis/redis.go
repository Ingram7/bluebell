package redis

import (
	"context"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: 100,
	})
	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}

func GetRdb() *redis.Client {
	return rdb
}
