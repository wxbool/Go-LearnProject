package main

import (
	"fmt"
	appclick "github.com.wxbool/go.learn.project/win.control.simulate/clicks/app"
	"log"
	"time"
)

//配置文件
const CONFIG = "click.ini"

//模拟点击小程序
func main() {

	//致命错误捕获
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error : %v", err)

			time.Sleep(time.Second * 8)
		}
	}()

	app := appclick.NewApp(CONFIG)

	loop := 0 //任务执行次数
	for loop < app.Loops {
		app.Run()

		loop++
	}

	fmt.Println("finish .")

	time.Sleep(time.Millisecond * 500)

}