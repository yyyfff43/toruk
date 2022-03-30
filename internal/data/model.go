/*
* @File : data
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/15 12:06
* @Software: GoLand
 */

package data

import (
	"git.zhwenxue.com/zhgo/gocontrib/db"
	"git.zhwenxue.com/zhgo/gocontrib/log"
	"git.zhwenxue.com/zhgo/toruk/internal/biz/dependency"
	"git.zhwenxue.com/zhgo/toruk/internal/pkg/resource"
	"xorm.io/xorm"
)



func NewModel(mysqlClient *xorm.Engine, redisClient *db.Redis, log *log.SugarLogger) (m dependency.Model, cf func(), err error){
	m = &model{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
		log: log,
	}
	cf = m.Close
	err = nil
	return
}

type model struct {
	mysqlClient *xorm.Engine
	redisClient *db.Redis
	log         *log.SugarLogger
}

func (m *model) Close() {
	_ = m.mysqlClient.Close()
	_ = m.redisClient.Close()
}

func NewTestModel() (dependency.Model, func()) {
	logger, err := resource.NewBasicLogger()
	if err != nil {
		panic(err)
	}
	engine, err := resource.NewMysqlClient(logger)
	if err != nil {
		panic(err)
	}
	redis,  err := resource.NewRedisClient(logger)
	if err != nil {
		panic(err)
	}
	sugarLogger,  err := resource.NewLogger()
	if err != nil {
		panic(err)
	}
	model, closeFunc, err := NewModel(engine, redis, sugarLogger)
	if err != nil {
		panic(err)
	}
	return model, closeFunc
}