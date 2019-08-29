package main

import (
	"fmt"
	appclick "github.com.wxbool/go.learn.project/win.control.simulate/clicks/app"
	"time"
)

//配置文件
const CONFIG = "click.ini"

//模拟点击小程序
func main() {

	app := appclick.NewApp(CONFIG)

	app.Run()

	fmt.Println("finish .")

	time.Sleep(time.Millisecond * 500)

}