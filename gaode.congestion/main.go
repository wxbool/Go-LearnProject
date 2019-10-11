package main

import (
	"./gaode"
	"log"
	"time"
)

//定义配置文件
const CONFIG = "config.ini"

func main() {

	//致命错误捕获
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error : %v", err)

			time.Sleep(time.Second * 8)
		}
	}()

	//获取应用
	app := gaode.NewApp(CONFIG)

	//初始化应用
	app.Init()

	//运行任务
	app.Run()

	//延迟退出
	time.Sleep(time.Second * 10)
}
