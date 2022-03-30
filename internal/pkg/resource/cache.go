package resource

import (
	"git.zhwenxue.com/zhgo/gocontrib/cache"
	Hlog "git.zhwenxue.com/zhgo/gocontrib/log"
	"git.zhwenxue.com/zhgo/toruk/internal/conf"
)

func NewCache(log Hlog.Logger) cache.Cache {
	configFile := conf.GetConfigFilePath("cache_config.yaml")
	cacheConfig, err := cache.ConfigWithPath(configFile)
	if err != nil {
		return nil
	}
	return cache.NewCache(cacheConfig, &log)
}
