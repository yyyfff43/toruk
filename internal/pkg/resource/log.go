/*
* @File : log
* @Describe :
* @Author: zhangnaiqian@zongheng.com
* @Date : 2022/1/17 14:30
* @Software: GoLand
 */

package resource

import (
	Hlog "git.zhwenxue.com/zhgo/gocontrib/log"
	"os"
)

func NewBasicLogger() (Hlog.Logger, error) {
	logger := Hlog.New(os.Stdout, Hlog.InfoLevel, Hlog.WithCaller(true), Hlog.AddCallerSkip(1))
	return *logger, nil
}

func NewLogger() (*Hlog.SugarLogger, error) {
	logger := Hlog.New(os.Stdout, Hlog.InfoLevel, Hlog.WithCaller(true), Hlog.AddCallerSkip(1))
	return logger.Sugar(), nil
}
