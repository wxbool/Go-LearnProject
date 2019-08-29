package main


func main() {


	//模拟热键WIN+R
	//robotgo.KeyTap("R", "command") //模拟热键WIN+R

	//KeyToggle的使用，按下或抬起按键
	//.KeyToggle(按键,按下或抬起，按键，按键...)
	//.KeyToggle(按键,"down"/"up"，按键，按键...)
	//robotgo.KeyToggle("r", "down" , "command")
	//robotgo.KeyToggle("a" , "down", "alt", "control") //长按ctrl+alt+a,如果打开了QQ，按下A可以截图




	//TypeStr的使用，在输入框中输入字符串
	//这个有缺陷，Hello真的打印出来可能是Helo，少字母
	//TypeStr(要打印的字符)
	//robotgo.TypeStr("Hello")

	//TypeStrDelay的使用，在输入框中输入字符串
	//TypeStrDelay(要打印的字符,每分钟输出的字符数)
	//robotgo.TypeStrDelay("Hello",500)




	//MoveMouse，移动鼠标到目标位置
	//MoveMouse(横向坐标,纵向坐标)
	//robotgo.MoveMouse(600,400)

	//MoveMouseSmooth，平滑移动到目标位置，模仿人类操作
	//MoveMouseSmooth(横向坐标,纵向坐标)
	//robotgo.MoveMouseSmooth(600,400)
	//robotgo.MoveMouseSmooth(600,400,1.0 , 20.0)//后面两个参数文档上看lowspeed, highspeed，与速度相关，是能改变移动速度，但是没搞清究竟是怎么个改变的

	//MouseClick，鼠标点击事件
	//MouseClick(哪一个键,是否双击)
	//MouseClick("left/center/right",true/false)
	//robotgo.MouseClick("left" , false)

	//MoveClick，移动到目标位置，并点击
	//MoveClick(横向坐标,纵向坐标,"right",false)
	//robotgo.MoveClick(1085,15,"left" , false)

	//MouseToggle,长按
	//MouseToggle("down/up","left/center/right")
	//robotgo.MouseToggle("down","left")

	//DragMouse名，拖动鼠标
	//实例里面与MouseToggle合用，没发现与MoveMouse的区别，使用方法相同
	//robotgo.MouseToggle("down")
	//robotgo.DragMouse(600, 400)

	//GetMousePos，获取鼠标当前位置
	//fmt.Println(robotgo.GetMousePos())

	//ScrollMouse，滚动鼠标滚轮
	//ScrollMouse(滚动大小,向上或向下滚动)
	//ScrollMouse(10,"up/down")
	//robotgo.ScrollMouse(1, "up")






	//GetPixelColor，获取目标位置颜色
	//fmt.Println(robotgo.GetPixelColor(600,400))//结果：2b2b2b

	//GetScreenSize，获取屏幕大小
	//fmt.Println(robotgo.GetScreenSize())//结果：1920 1080，宽度   高度





	//位图操作**********************************************

	//CaptureScreen，截图操作
	//返回值为位图的对象
	//CaptureScreen(横向坐标x,纵向坐标y,图片高度,图片宽度)
	//截图的时候是在x,y点的基础分别加上高度与宽度的点，两点形成的矩形框
	//bitmap := robotgo.CaptureScreen(10, 20, 30, 40)

	//SaveBitmap，保存位图为文件
	//截了图，当然要保存才有用，这里的bitmap就是上面得到的bitmap
	//SaveBitmap(位图的对象, 保存路径与文件名,图片类型（可忽略，我也没搞清楚）)
	//robotgo.SaveBitmap(bitmap, "test.png")

	//OpenBitmap，把png图片转为bitmap对象
	//bitmap := robotgo.OpenBitmap("test.png")

	//FindBitmap，找到bitmap对象的坐标
	//fx, fy := robotgo.FindBitmap(bitmap)
	//fmt.Println("FindBitmap------", fx, fy)

	//TostringBitmap，把位图对象转为字符串
	//fmt.Println(robotgo.TostringBitmap(bitmap))

	//GetPortion，截取位图对象的一部分
	//GetPortion(位图对象,起始横向坐标,起始纵向坐标,宽度,高度)
	/**
		bm := robotgo.OpenBitmap("test.png")
		cutBm := robotgo.GetPortion(bm,0,0,100,100)
		robotgo.SaveBitmap(cutBm,"cutBm.png")
	*/

	//Convert，图片格式转换
	//Convert(文件源路径,新文件目标路径,文件格式)，文件格式文档没写，代码里也没看出对应的判断代码，不知道有什么意义，为0时大小不变，为2时文件大小变大很多，其他无作用，文件格式设置可省略
	//robotgo.Convert("test.png", "test.jpg",2)


	//FreeBitmap，释放位图
	//把位图释放掉，下面代码在保存时会报错
	/**
		bm := robotgo.OpenBitmap("test.png")
		robotgo.FreeBitmap(bm)
		robotgo.SaveBitmap(bm,"freeBm.png")
	 */


	//ReadBitmap，判断位图是否为空
	//如果为空，返回false,否则true
	/*
		bm := robotgo.OpenBitmap("test.png")
		bm=nil
		fmt.Println(robotgo.ReadBitmap(bm))
	 */


	//CopyBitpb，将位图复制到剪贴板
	//在word文件中ctrl+v，图片就会拷贝到文件中了
	/*
		bm := robotgo.OpenBitmap("test.png")
		robotgo.CopyBitpb(bm)
	 */

	//DeepCopyBit，复制位图对象
	/*
		bm := robotgo.OpenBitmap("test.png")
		cyBm := robotgo.DeepCopyBit(bm)
		robotgo.SaveBitmap(cyBm,"cyBm.png")
	 */





	//事件操作**********************************************

	//AddEvent，添加全局事件监听
	//AddEvent(按键或者鼠标操作)
	//支持的按键有数字字母等按键，支持的功能键从代码中可以看到有截图中的按键
	//鼠标(鼠标参数: mleft, mright, wheelDown, wheelUp, wheelLeft, wheelRight)，目前测试中只有前面两个对应代码有效
	/*
		lEvent := robotgo.AddEvent("mleft")
		if lEvent == 0 {
			fmt.Println("mleft")
		}
	 */

	//StopEvent()，停止事件监听
	//暂时没想到用到的地方，测试时要把监听也停止监听跑两个线程才能看到效果，AddEvent运行时是阻塞的
	//robotgo.StopEvent()





	//窗口操作**********************************************
	//几个文档中的函数测试中无法有效使用（可能是我win10系统），或者没找到合适用法的，没有写入下方文档

	//ShowAlert，消息提示框
	//ShowAlert("标题", "消息", "Success","Close")，后面两个参数无效，可省略，中文存在乱码问题
	/*
		btMsg := robotgo.ShowAlert("Title", "This Message!", "Success","Close")
		fmt.Println(btMsg)//确定0，取消1
	 */


	//GetActive，SetActive获取当前窗口和跳转到某窗口
	//下面代码测试流程：程序正式执行后，5S类切换到另外一个窗口，过会又会跳到开始的窗口
	/*
		cw := robotgo.GetActive()//获取当前选择的窗口
		time.Sleep(5*time.Second)
		robotgo.SetActive(cw)//跳转到目标窗口
	 */


	//GetTitle，获取窗口标题
	//fmt.Println(robotgo.GetTitle())


	//GetHandle，SetHandle，获取窗口句柄和设置句柄
	//获取没问题，但设置发现无效
	/*
		fmt.Println(robotgo.GetHandle())
		robotgo.SetHandle(3272727)
	 */


	//Pids，获取进程PID
	//返回两个参数，第一个是pid的数组，第二个是错误信息
	//fmt.Println(robotgo.Pids())


	//PidExists，判断进程是否存在
	//根据PID判断，返回两个参数，第一个bool值存在true，第二个是错误信息
	//fmt.Println(robotgo.PidExists(928))


	//Process，获取进程信息
	//返回两个参数：进程信息的数组，错误信息
	//fmt.Println(robotgo.Process())


	//FindName，根据PID查询进程名
	//返回参数：进程名，错误信息
	//fmt.Println(robotgo.FindName(928))


	//FindNames，查询所有进程名
	//不知道是不是系统原因，无效
	//返回参数：进程名数组，错误信息
	//fmt.Println(robotgo.FindNames())


	//FindIds，根据进程名找PID
	//返回参数：进程PID，错误信息
	//fmt.Println(robotgo.FindIds("lsass.exe"))


	//ActivePID，根据PID激活窗口
	//不知道是不是系统原因，无效
	//返回参数：错误信息
	//fmt.Println(robotgo.ActivePID(9792))
}
