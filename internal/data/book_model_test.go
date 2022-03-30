/*
* @File : book_model_test
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/18 15:32
* @Software: GoLand
 */

package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBookInfo(t *testing.T) {
	model, closeFunc := NewTestModel()
	defer func() {
		closeFunc()
	}()
	data, err := model.GetBookInfo(Context, 1)
	assert.Nil(t, err)
	fmt.Println(data)
}