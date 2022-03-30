/*
* @File : biz_test
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/18 19:54
* @Software: GoLand
 */

package biz

import (
	"context"
	"git.zhwenxue.com/zhgo/dubbogorpc/ugc"
	zhContext "git.zhwenxue.com/zhgo/gocontrib/context"
	"git.zhwenxue.com/zhgo/toruk/internal/data"
	"git.zhwenxue.com/zhgo/toruk/internal/pkg/resource"
	"git.zhwenxue.com/zhgo/toruk/internal/pkg/rpc"
)

var Context = zhContext.GetContext(context.Background(), "testuuid")

func NewTestBiz() (Biz, func()) {
	model, closeModel := data.NewTestModel()

	logger, err := resource.NewLogger()
	if err != nil {
		panic(err)
	}
	biz := NewBiz(model, logger, NewUgcRpcClient())
	return biz, closeModel
}

func NewUgcRpcClient() *ugc.FavoriteBookService{
	return rpc.NewUgcService()
}