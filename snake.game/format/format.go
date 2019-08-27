package format

import (
	"math/rand"
	"strconv"
	"time"
)

//行进方向
const (
	PATH_UP = 1
	PATH_DOWN = 2
	PATH_LEFT = 3
	PATH_RIGHT = 4
)

type Food struct {
	x int //x坐标
	y int //y坐标
}

type GameCanvas struct {
	Width int
	Height int

	GameFood Food //食物坐标
	FoodSymbol string //食物符号
	Grade int //分数
	ReduceSpeed time.Duration //蛇移动速度递减量
}

type Snake struct { //一条蛇
	len int //总长度
	maxX int //最大x坐标
	maxY int //最大y坐标
	body [][]int //蛇坐标
	path int //行进方向
	speed time.Duration //速度
	symbol string //蛇体符号
	fail string //失败信息
}


//创建新面板
func (canvas *GameCanvas) NewGameCanvas(snake Snake) string {
	//随机生成一个食物
	fx , fy := RandomFood(snake.maxX , snake.maxY)
	canvas.GameFood.x = fx
	canvas.GameFood.y = fy
	canvas.Grade = 0

	return drawCanvas(*canvas , snake)
}

//创建一条蛇
func (canvas GameCanvas) NewSnake(len int , body [][]int , symbol string , speed time.Duration) Snake {
	snake := Snake{
		len:len,
		body:body,
		maxX:canvas.Width-3,
		maxY:canvas.Height-3,
		speed:speed,
		symbol:symbol,
	}
	return snake
}

//改变蛇的方向
func (canvas GameCanvas) ChangeSnakePath(snake *Snake , path int) {
	switch snake.path {
		case PATH_UP:
			if path == PATH_DOWN {
				return
			}
		case PATH_DOWN:
			if path == PATH_UP {
				return
			}
		case PATH_LEFT:
			if path == PATH_RIGHT {
				return
			}
		case PATH_RIGHT:
			if path == PATH_LEFT {
				return
			}
	}
	snake.path = path
}

//将蛇往下移
func (canvas *GameCanvas) MoveDownSnake(snake *Snake) string {
	newhead := []int{snake.body[0][0] , snake.body[0][1]+1} //设置头部坐标
	replaceSnake(newhead , snake , canvas)
	return drawCanvas(*canvas , *snake)
}

//将蛇往上移
func (canvas *GameCanvas) MoveUpSnake(snake *Snake) string {
	newhead := []int{snake.body[0][0] , snake.body[0][1]-1} //设置头部坐标
	replaceSnake(newhead , snake , canvas)
	return drawCanvas(*canvas , *snake)
}

//将蛇向右移动
func (canvas *GameCanvas) MoveRightSnake(snake *Snake) string {
	newhead := []int{snake.body[0][0]+1 , snake.body[0][1]}
	replaceSnake(newhead , snake , canvas)
	return drawCanvas(*canvas , *snake)
}

//将蛇向左移动
func (canvas *GameCanvas) MoveLeftSnake(snake *Snake) string {
	newhead := []int{snake.body[0][0]-1 , snake.body[0][1]} //设置头部坐标
	replaceSnake(newhead , snake , canvas)
	return drawCanvas(*canvas , *snake)
}

//随机获取一个食物坐标
func RandomFood(maxX int , maxY int) (x int , y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(maxX)
	rand.Seed(time.Now().UnixNano())
	y = rand.Intn(maxY)
	return
}

//替换蛇身体
func replaceSnake(newhead []int , snake *Snake , canvas *GameCanvas) {
	len := len(snake.body)
	newbody := make([][]int , 0 , len)

	//吃到食物
	if (newhead[0] == canvas.GameFood.x) && (newhead[1] == canvas.GameFood.y) {
		//吃到果子，追加蛇节点
		newbody = append(newbody , []int{newhead[0] , newhead[1]})
		for index , _ := range snake.body {
			newbody = append(newbody , []int{snake.body[index][0] , snake.body[index][1]})
		}
		//重新生成食物
		fx , fy := RandomFood(snake.maxX , snake.maxY)
		canvas.GameFood.x = fx
		canvas.GameFood.y = fy
		//速度递减
		snake.speed -= canvas.ReduceSpeed
		//加分数
		canvas.Grade++
	} else  {
		//移动蛇节点
		for index , _ := range snake.body {
			if index == 0 {
				newbody = append(newbody , []int{newhead[0] , newhead[1]})
			} else {
				newbody = append(newbody , []int{snake.body[index - 1][0] , snake.body[index - 1][1]})
			}
		}
	}

	//检测蛇身冲突
	//检测蛇是否撞墙
	for index , _ := range newbody {
		if index != 0 {
			if (newhead[0] == newbody[index][0]) && (newhead[1] == newbody[index][1]) {
				//冲突
				snake.fail = "You lose";
			}
		}

		if newbody[index][0] > snake.maxX || newbody[index][0] < 0 {
			//撞墙
			snake.fail = "You lose";
		}
		if newbody[index][1] > snake.maxY || newbody[index][1] < 0 {
			//撞墙
			snake.fail = "You lose";
		}
	}

	snake.body = newbody
}


//蛇自动行进
func (canvas GameCanvas) Watch(snake *Snake , callback func (string))  {
	var box string

	for {

		time.Sleep(time.Millisecond * snake.speed)
		if snake.path == 0 {
			continue
		}

		//蛇行进失败
		if snake.fail != "" {
			box := drawCanvas(canvas , *snake)
			box += "\n" + snake.fail
			callback(box)
			break
		}

		switch snake.path {
			case PATH_UP:
				box = canvas.MoveUpSnake(snake)
			case PATH_DOWN:
				box = canvas.MoveDownSnake(snake)
			case PATH_LEFT:
				box = canvas.MoveLeftSnake(snake)
			case PATH_RIGHT:
				box = canvas.MoveRightSnake(snake)
		}

		callback(box)
	}
}


//输出游戏面板
func drawCanvas(canvas GameCanvas , snake Snake) string {
	fmtout := ""
	fmtout += "\n\n\n\n\n"

	for h:=0;h<canvas.Height;h++ {
		fmtout += "       " //间隔

		if h == 0 {
			fmtout += "┌"
		} else if (h == canvas.Height-1) {
			fmtout += "└"
		} else {
			fmtout += "┆"
		}

		for w:=0;w<canvas.Width;w++ {
			snakeBool := false
			//注入蛇
			if snake.len != 0 {
				for _ , val := range snake.body {
					if val[0] == getX(canvas , w) && val[1] == getX(canvas , h) {
						fmtout += snake.symbol
						snakeBool = true
					}
				}
			}
			//注入食物
			if canvas.GameFood.x == getX(canvas , w) && canvas.GameFood.y == getX(canvas , h) {
				fmtout += canvas.FoodSymbol
				snakeBool = true
			}
			if snakeBool {
				continue
			}

			if w==canvas.Width-1 {
				if h == 0 {
					fmtout += "┐"
				} else if h == canvas.Height - 1 {
					fmtout += "┘"
				} else {
					fmtout += "┆"
				}
			} else {
				if (h == 0) || (h == canvas.Height - 1) {
					fmtout += "┈"
				} else {
					fmtout += " "
				}
			}
		}
		fmtout += "\n"
	}

	//输出得分
	fmtout += "\nYour score : " + strconv.Itoa(canvas.Grade)

	//返回面板数据
	return fmtout
}


//画布x坐标 -> 蛇x坐标
func getX(canvas GameCanvas , x int) int {
	if x == 0 || x == canvas.Width-1 {
		return -1
	}
	return x-1
}

//画布y坐标 -> 蛇y坐标
func getY(canvas GameCanvas , y int) int {
	if y == 0 || y == canvas.Height-1 {
		return -1
	}
	return y-1
}