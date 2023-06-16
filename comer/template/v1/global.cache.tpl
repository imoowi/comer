package global

import (
	"context"
	"strings"
	"time"

	format "github.com/imoowi/comer/util/format"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

func getCachePrefix() string {
	return Config.GetString(`settings.cache.prefix`) + ":"
}
func GetCache(key string) (data string, err error) {
	cacheKey := getCachePrefix() + key
	data, err = Redis.Get(context.Background(), cacheKey).Result()
	if err != nil {
		// fmt.Println(`cache get err:`, err)
		return
	}
	return
}

func SetCache(key string, data interface{}, duration time.Duration) bool {
	cacheKey := getCachePrefix() + key
	err := Redis.Set(context.Background(), cacheKey, data, duration).Err()
	// fmt.Println(`cache set err:`, err)
	return err == nil
}

// 批量获取缓存
func GetCacheBatch(keys ...string) (map[string]string, error) {
	keys = format.UniqueSliceString(keys)
	keysFormat := make([]string, len(keys))
	for idx, key := range keys {
		keysFormat[idx] = getCachePrefix() + key
	}

	res, err := Redis.MGet(context.Background(), keysFormat...).Result()
	if err != nil {
		Log.Logger.WithFields(logrus.Fields{
			"keys": keys,
		}).Errorf("get cache batch error: %s\n", err)
		return nil, nil
	}
	cachePrefix := getCachePrefix()
	data := make(map[string]string)
	for idx, key := range keysFormat {
		if res[idx] != nil {
			key := strings.TrimPrefix(key, cachePrefix)
			data[key] = cast.ToString(res[idx])
		}
	}
	return data, nil
}
