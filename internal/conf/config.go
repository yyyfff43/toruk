/*
* @File : config
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/17 10:31
* @Software: GoLand
 */

package conf

import (
	"git.zhwenxue.com/zhgo/toruk/internal/common/constant"
	"os"
)

func GetConfigFilePath(fileName string) string {
	dir, _ := os.Getwd()
	configFilePath := dir + "/internal/conf/dev/" + fileName
	if configFilePathFromEnv := os.Getenv(constant.APPConfigPath); configFilePathFromEnv != "" {
		configFilePath = configFilePathFromEnv + "/" + fileName
	}

	return configFilePath
}
