//+build wireinject


package assembly

import (
	"git.zhwenxue.com/zhgo/toruk/internal/biz"
	"git.zhwenxue.com/zhgo/toruk/internal/data"
	"git.zhwenxue.com/zhgo/toruk/internal/pkg/resource"
	"git.zhwenxue.com/zhgo/toruk/internal/pkg/rpc"
	dubbo2 "git.zhwenxue.com/zhgo/toruk/internal/server/dubbo"
	"git.zhwenxue.com/zhgo/toruk/internal/service/dubbo"
	"github.com/google/wire"
)

func NewUserDubboServer() (dubbo2.DubboServer, func(), error) {
	panic(wire.Build(
		resource.ProviderSet,
		data.NewModel,
		rpc.NewUgcService,
		biz.NewBiz,
		dubbo.NewUserProvider,
		dubbo2.NewUserDubboServer,
	))
}