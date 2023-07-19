package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Cache *CacheRedis

type CacheRedis struct {
	client *redis.Client
}

// SetKeyValue 设置
func (c *CacheRedis) SetKeyValue(prefix string, key any, value any, expiration time.Duration) (keyStr string, err error) {
	keyStr = fmt.Sprintf("%s:%v", prefix, key)
	err = c.client.Set(keyStr, value, expiration).Err()
	return
}

// GetKey 获取
func (c *CacheRedis) GetKey(prefix string, key any) (string, error) {
	keyStr := fmt.Sprintf("%s:%v", prefix, key)
	return c.client.Get(keyStr).Result()
}

func (c *CacheRedis) Delete(prefix string, key any) error {
	keyStr := fmt.Sprintf("%s:%v", prefix, key)
	return c.client.Del(keyStr).Err()
}

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
	Log.Info("连接Redis数据库成功")
	Cache = &CacheRedis{client: client}
}
