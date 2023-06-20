/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package global

import (
	"log"

	"github.com/go-redis/redis/v8"
	"{{.moduleName}}/components"
)

// 全局Redis客户端
var Redis *redis.Client

// 初始化redis
func initRedis() {
	// 获取logger相关的配置信息
	config := Config.Sub("redis")
	var redisConfig *components.RedisConfig
	err := config.Unmarshal(&redisConfig)
	if err != nil {
		log.Fatal(err)
	}
	redisClient := components.NewRedisClient(redisConfig)
	Redis = redisClient
}
