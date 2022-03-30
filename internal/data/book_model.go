/*
* @File : book
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/15 12:04
* @Software: GoLand
 */

package data

import (
	"context"
	"fmt"
	"git.zhwenxue.com/zhgo/toruk/internal/entity"
	"time"
	"xorm.io/builder"
)



func (m *model) GetBookInfo(ctx context.Context, id int) (*entity.Book, error) {
	book := &entity.Book{}
	where := builder.Eq{}
	where["book_id"] = id
	_, err := m.mysqlClient.Context(ctx).Where(where).Get(book)
	m.redisClient.Set(ctx, "test0", "123", 5 * time.Second)
	cmd := m.redisClient.Get(ctx, "test0")
	fmt.Println(cmd.Val())
	return book, err
}