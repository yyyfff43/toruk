/*
* @File : BookModel
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/16 11:31
* @Software: GoLand
 */

package dependency

import (
	"context"
	"git.zhwenxue.com/zhgo/toruk/internal/entity"
)

type BookModel interface {
	GetBookInfo(ctx context.Context, id int) (*entity.Book, error)
}