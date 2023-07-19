package core

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Cache *redis.Client

func InitRedis() {
	config := Config.Redis
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.Db,
	})
	if _, err := client.Ping().Result(); err != nil {
		Log.Error("连接Redis数据库失败: %s", err.Error())
		panic(err.Error())
	}
	Log.Error("连接Redis数据库成功")
	Cache = client
}
