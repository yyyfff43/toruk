/*
* @File : ModelInterface
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/16 11:36
* @Software: GoLand
 */

package dependency

type Model interface {
	BookModel
	UserModel
	Close()
}
