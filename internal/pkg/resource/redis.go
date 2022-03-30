/*
* @File : redis
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/17 10:02
* @Software: GoLand
*/

package resource

import (
	Hdb "git.zhwenxue.com/zhgo/gocontrib/db"
	Hlog "git.zhwenxue.com/zhgo/gocontrib/log"
	"git.zhwenxue.com/zhgo/toruk/internal/conf"
)

func NewRedisClient(log Hlog.Logger) (*Hdb.Redis, error){
	configFile := conf.GetConfigFilePath("redis_config.yaml")
	cfg, err := Hdb.RedisConfigWithPath(configFile)
	if err != nil {
		panic(err)
	}
	redis, err := Hdb.NewRedis(cfg, log)
	if err != nil {
		panic(err)
	}
	return redis, err
}

//func LoadConfiguration(path string) (*Hdb.Config, error) {
//	var cfg *Hdb.Config
//	data, err := ioutil.ReadFile(path)
//	if err != nil {
//		return cfg, err
//	}
//	err = yaml.Unmarshal(data, &cfg)
//	if err != nil {
//		return cfg, err
//	}
//	return cfg, nil
//}
