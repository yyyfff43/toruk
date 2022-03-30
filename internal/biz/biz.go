/*
* @File : biz
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/16 2:14
* @Software: GoLand
 */

package biz

import (
	"context"
	"git.zhwenxue.com/zhgo/dubbogorpc/ugc"
	"git.zhwenxue.com/zhgo/gocontrib/cache"
	"git.zhwenxue.com/zhgo/gocontrib/log"
	"git.zhwenxue.com/zhgo/toruk/internal/biz/dependency"
	"git.zhwenxue.com/zhgo/toruk/internal/entity"
)

type Biz interface {
	GetUserFavoriteBook(ctx context.Context, userID int) (book *entity.Book, err error)
}

type biz struct {
	model dependency.Model
	log   *log.SugarLogger
	ugc   *ugc.FavoriteBookService
	cache cache.Cache
}

func NewBiz(model dependency.Model, log *log.SugarLogger, ugc *ugc.FavoriteBookService, cache cache.Cache) Biz {
	return &biz{model: model, log: log, ugc: ugc, cache: cache}
}
