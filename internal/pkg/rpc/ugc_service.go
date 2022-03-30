/*
* @File : ugc_service
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/17 16:40
* @Software: GoLand
 */

package rpc

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"git.zhwenxue.com/zhgo/dubbogorpc/ugc"
	"git.zhwenxue.com/zhgo/toruk/internal/conf"
)



func NewUgcService() *ugc.FavoriteBookService{
	configFile := conf.GetConfigFilePath("ugc_rpc.yml")
	if err := config.Load(config.WithPath(configFile)); err != nil {
		panic(err)
	}

	return ugc.FavoriteBookClient
}