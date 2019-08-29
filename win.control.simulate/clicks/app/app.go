package app

import (
	"github.com.wxbool/go.learn.project/config.file/ini"
	"github.com/go-vgo/robotgo"
	"strconv"
	"strings"
	"time"
)

type Coordinate struct {
	X int //X坐标点
	Y int //Y坐标点
	Click string //点击鼠标按键 => left/center/right
	Dbclick bool //是否双击
	Nums int //点击次数
}

//点击器配置
type ClickCfg struct {
	Coordinate []*Coordinate //点击坐标点切片
	StartWait int //任务开始前等待时间（秒s）
	ClickWait int //多个坐标点点击间隔时间（毫秒ms），1000ms = 1s
}


//根据配置文件实例应用
func NewApp(cfg string) (app *ClickCfg) {
	//读取配置
	app = readConfig(cfg)
	return
}


//运行
func (app *ClickCfg) Run() {
	//开始运行等待(s)
	time.Sleep(time.Second * time.Duration(app.StartWait))

	//遍历点击
	for _ , item := range app.Coordinate {
		if item.Nums <= 0 {
			continue
		}
		//执行次数
		index := 0
		for index < item.Nums {
			//移动到目标坐标
			robotgo.MoveMouseSmooth(item.X , item.Y)
			//鼠标点击
			robotgo.MouseClick(item.Click  , item.Dbclick)

			//执行后间隔等待
			time.Sleep(time.Millisecond * time.Duration(app.ClickWait))

			index++
		}
	}
}


//获取配置
func readConfig(cfg string) *ClickCfg {
	if c , err := ini.LoadConfigFile(cfg , ".");err != nil {
		panic(err)
	} else {
		clickCfg := &ClickCfg{}
		clickCfg.StartWait = c.GetIntMust("task.startwait" , 0)
		clickCfg.ClickWait = c.GetIntMust("task.clickwait" , 0)

		//获取坐标列表
		coordinate := c.GetMust("task.coordinate" , "")
		//解析坐标列表
		clickCfg.Coordinate = readCoordinate(coordinate)

		return clickCfg
	}
}

//解析坐标列表
func readCoordinate(coord string) []*Coordinate {
	if coord == "" {
		panic("Task coordinate points is empty .")
	}

	taskSlice := make([]*Coordinate , 0 , 20)
	//分隔
	coordSlice := strings.Split(coord , "|")

	for _ , co := range coordSlice {
		thisCoord := strings.Split(co , ",")
		//检测格式
		if len(thisCoord) != 5 {
			panic("Task coordinate point input format error .")
		}

		x := thisCoord[0]
		y := thisCoord[1]
		click := thisCoord[2]
		dbclick := thisCoord[3]
		nums := thisCoord[4]

		xVal , err := strconv.Atoi(x)
		if err != nil {
			panic("X coordinate error .")
		}
		yVal , err := strconv.Atoi(y)
		if err != nil {
			panic("X coordinate error .")
		}
		//left/center/right
		if click != "left" && click != "center" && click != "right" {
			panic("coordinate error : "+ click)
		}
		dbVal , err := strconv.ParseBool(dbclick)
		if err != nil {
			panic("coordinate error : " + dbclick)
		}
		numsVal , err := strconv.Atoi(nums)
		if err != nil {
			panic("coordinate error : " + nums)
		}

		taskSlice = append(taskSlice , &Coordinate{X:xVal , Y:yVal , Click:click , Dbclick:dbVal , Nums:numsVal})
	}
	return taskSlice
}