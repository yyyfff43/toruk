/*
* @File : userService
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/16 3:02
* @Software: GoLand
 */

package dubbo

import (
	"context"
	"git.zhwenxue.com/zhgo/toruk/internal/biz"
)

type UserProvider interface {
	GetUserFavoriteBook(ctx context.Context, id int64)
}

func NewUserProvider (biz biz.Biz) UserProvider {
	return &userProvider{biz: biz}
}

type userProvider struct {
	biz biz.Biz
}

func (u *userProvider) GetUserFavoriteBook(ctx context.Context, id int64)  {

}