/*
* @File : user
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/16 2:10
* @Software: GoLand
 */

package biz

import (
	"context"
	"fmt"
	"git.zhwenxue.com/zhgo/toruk/internal/entity"
)

func (b *biz) GetUserFavoriteBook(ctx context.Context, userID int) (book *entity.Book, err error) {
	bookID, err := b.model.GetUserPreference(ctx, userID)
	//b.log.Sugar().Info()
	book, err = b.model.GetBookInfo(ctx, bookID)
	books, err := b.ugc.GetFavoriteBooks(ctx, 1, 0, 1, 20)
	fmt.Println(books)
	return book, err
}