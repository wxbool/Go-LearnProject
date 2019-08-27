package app

import "github.com/nsf/termbox-go"

type AppRun struct {
	UpEvent func()
	DownEvent func()
	LeftEvent func()
	RightEvent func()
	EnterEvent func()
	EscEvent func()
}

//开始运行
func (app AppRun) Run() {

	//注册监听
	//初始化
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	//结束的时候,记得关闭
	defer termbox.Close()

	//阻塞监听键盘事件
Loop:
	for {
		//termbox.PollEvent()  这个监控事件输入,保存到ev
		//ev.Type   事件类型
		//ev.Key    一些特殊按键,比如  key+c F1 F2 Esc Enter等
		//ev.Ch       正常的字符健.  abcd 1234
		//termbox.EventKey   判断是否键盘事件
		//termbox.EventResize  判断是否调整窗口大小
		//termbox.EventMouse   判断是否鼠标事件
		//termbox.EventError   判断错误事件
		//termbox.EventInterrupt  判断终止事件.

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case 27: //Esc
				app.EscEvent() //退出事件
				break Loop
			case 65517: //Up
				app.UpEvent()
			case 65516: //Down
				app.DownEvent()
			case 65515: //Left
				app.LeftEvent()
			case 65514: //Right
				app.RightEvent()
			case 13: //Enter
				app.EnterEvent()
			}

		}
	}
}