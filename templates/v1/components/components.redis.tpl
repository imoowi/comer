/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package components

import "github.com/go-redis/redis/v8"

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
