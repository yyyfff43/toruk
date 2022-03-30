/*
* @File : user_test
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/18 19:54
* @Software: GoLand
 */

package biz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGetUserFavoriteBook(t *testing.T) {
	biz, clean := NewTestBiz()
	defer func() {
		clean()
	}()
	book, err := biz.GetUserFavoriteBook(Context, 1)
	assert.Nil(t, err)
	fmt.Println(book)
}
