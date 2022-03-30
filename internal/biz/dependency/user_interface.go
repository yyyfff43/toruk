/*
* @File : UserInterface
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/16 11:33
* @Software: GoLand
 */

package dependency

import "context"

type UserModel interface {
	CheckUser(ctx context.Context, id int) (error)
	GetUserPreference(ctx context.Context, id int) (int, error)
}