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
func (c *CacheRedis) SetKeyValue(key string, value any, expiration time.Duration) (keyStr string, err error) {
	err = c.client.Set(key, value, expiration).Err()
	return
}

// GetKey 获取
func (c *CacheRedis) GetKey(key string) (string, error) {
	return c.client.Get(key).Result()
}

// Delete 删除
func (c *CacheRedis) Delete(key ...string) error {
	return c.client.Del(key...).Err()
}

// Exist 判断key存在
func (c *CacheRedis) Exist(key string) bool {
	return c.client.Exists(key).Val() == 1
}

// IsExpire 获取key的过期时间
// 当 key 不存在时，返回 -2 。
// 当 key 存在但没有设置剩余生存时间时，返回 -1
// 以秒为单位，返回 key 的剩余生存时间
func (c *CacheRedis) IsExpire(key string) float64 {
	return c.client.TTL(key).Val().Seconds()
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
