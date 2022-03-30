/*
* @File : mysql
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/17 10:03
* @Software: GoLand
 */

package resource

import (
	"git.zhwenxue.com/zhgo/gocontrib/db"
	Hlog "git.zhwenxue.com/zhgo/gocontrib/log"
	"git.zhwenxue.com/zhgo/toruk/internal/conf"
	"xorm.io/xorm"
)

func NewMysqlClient(log Hlog.Logger) (*xorm.Engine, error){
	configFile := conf.GetConfigFilePath("mysql_config.yaml")
	mysqlConfig, err :=  db.MysqlConfigWithPath(configFile)
	if err != nil {
		return nil, err
	}
	return db.NewMysqlClient(mysqlConfig, log)
}