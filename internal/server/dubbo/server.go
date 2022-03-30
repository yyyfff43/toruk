/*
* @File : server
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/17 1:13
* @Software: GoLand
 */

package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
	"git.zhwenxue.com/zhgo/toruk/internal/service/dubbo"
	"os"
)

type DubboServer interface {
	Start() error
}

type userDubboServer struct {
	u dubbo.UserProvider
}

func NewUserDubboServer(u dubbo.UserProvider) DubboServer {
	return &userDubboServer{u: u}
}

func (u *userDubboServer) Start() error {
	config.SetProviderService(u) // 注册服务提供者类，类名与配置文件中的 service 对应
	dir, _ := os.Getwd()
	configFile := dir + "/internal/server/dubbo/dubbo.yml"
	if err := config.Load(config.WithPath(configFile)); err != nil {
		return err
	}
	return nil
}