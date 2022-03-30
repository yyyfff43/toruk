/*
* @File : user
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/16 0:42
* @Software: GoLand
 */

package data

import (
	"context"
)

func (m *model) GetUserPreference(ctx context.Context, id int) (int, error) {
	return id, nil
}

func (m *model) CheckUser(ctx context.Context, id int) error {
	return nil
}