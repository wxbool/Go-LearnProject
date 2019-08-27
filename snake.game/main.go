package main

import (
	"fmt"
	"github.com/gosuri/uilive"
	appGame "snake.game/app"
	"snake.game/format"
	"time"
)

func main() {

	var app = appGame.AppRun{}
	var game = format.GameCanvas{Width:50,Height:25,FoodSymbol:"$",ReduceSpeed:5} //创建实例
	var snake = game.NewSnake(5 , [][]int{{3,3} , {3,4} , {3,5}} , "O" , 180) //初始化蛇
	var init = false //游戏初始化

	//缓冲区
	writer := uilive.New()
	writer.Start()

	//注册事件
	app.EnterEvent = func() {
		if !init {
			box := game.NewGameCanvas(snake)
			fmt.Fprintln(writer, box)

			//自动行进注册
			go func() {
				game.Watch(&snake , func(box string) {
					fmt.Fprintln(writer, box)
				})
			}()
		}
		init = true
	}
	//上移
	app.UpEvent = func() {
		game.ChangeSnakePath(&snake , format.PATH_UP)
	}
	//下移
	app.DownEvent = func() {
		game.ChangeSnakePath(&snake , format.PATH_DOWN)
	}
	//右移
	app.RightEvent = func() {
		game.ChangeSnakePath(&snake , format.PATH_RIGHT)
	}
	//左移
	app.LeftEvent = func() {
		game.ChangeSnakePath(&snake , format.PATH_LEFT)
	}

	//退出
	app.EscEvent = func() {
		fmt.Fprintln(writer, "游戏已退出")
	}

	fmt.Fprintln(writer, "按 Enter 键开始游戏")
	fmt.Fprintln(writer, "按 Esc 退出游戏")
	fmt.Fprintln(writer, "使用 上下左右键 来控制游戏")

	time.Sleep(time.Millisecond * 1500)

	app.Run()

	writer.Stop()
}

