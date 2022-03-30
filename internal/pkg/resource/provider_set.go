/*
* @File : provider_set
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/17 14:23
* @Software: GoLand
 */

package resource

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewBasicLogger, NewMysqlClient, NewRedisClient, NewLogger, NewCache)