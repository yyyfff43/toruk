package main

import (
	"git.zhwenxue.com/zhgo/toruk/assembly"
)

func main() {
	userServer, clean, err := assembly.NewUserDubboServer()
	if err != nil {
		if clean != nil {
			clean()
		}
		panic(err)
	}
	err = userServer.Start()
	panic(err)
}
