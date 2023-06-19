/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package global

import (
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string `json:"addr"`     // 地址
	Password string `json:"password"` // 密码
	DB       int    `json:"db"`       // 数据库
	PoolSize int    `json:"poolSize"` // Maximum number of socket connections.
}

func NewRedisClient(config *RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
	return rdb
}

// 全局Redis客户端
var Redis *redis.Client

// 初始化redis
func initRedis() {
	// 获取logger相关的配置信息
	config := Config.Sub("redis")
	var redisConfig *RedisConfig
	err := config.Unmarshal(&redisConfig)
	if err != nil {
		log.Fatal(err)
	}
	redisClient := NewRedisClient(redisConfig)
	Redis = redisClient
}
