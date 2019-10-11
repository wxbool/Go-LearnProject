package gaode

import (
	"./config/ini"
	"./model"
	"./datetime"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
)

//主应用
type AppGaode struct {
	//数据库连接配置
	DbConfig struct {
		Hostname string
		Database string
		Username string
		Password string
		Charset string
		Port string
		Prefix string
	}
	//数据库连接
	Db *gorm.DB

	Startdate string //运行开始时间
	Stopdate string //运行截止时间
	Ranks int //仅获取前 n 名城市排名的数据
	Intervals int   //获取间隔时间（s）秒
}


//获取应用
func NewApp(cfg string) *AppGaode {
	app := ReadConfig(cfg)

	return app
}


//读取配置
func ReadConfig (cfg string) *AppGaode {
	if file, e := ini.LoadConfigFile(cfg , ".");e != nil  {
		panic(e);
	} else {
		appconfig := &AppGaode{}
		appconfig.DbConfig.Hostname = file.GetMust("database.hostname" , "")
		appconfig.DbConfig.Database = file.GetMust("database.database" , "")
		appconfig.DbConfig.Username = file.GetMust("database.username" , "")
		appconfig.DbConfig.Password = file.GetMust("database.password" , "")
		appconfig.DbConfig.Charset = file.GetMust("database.charset" , "")
		appconfig.DbConfig.Prefix = file.GetMust("database.prefix" , "")
		appconfig.DbConfig.Port = file.GetMust("database.port" , "")

		appconfig.Startdate = file.GetMust("task.startdate" , "")
		appconfig.Stopdate = file.GetMust("task.stopdate" , "")
		appconfig.Ranks = file.GetIntMust("task.ranks" , 100)
		appconfig.Intervals = file.GetIntMust("task.intervals" , 300)

		return appconfig
	}
}


//日志输出
func Log(agrs ...interface{}) {
	fmt.Println(agrs ...)
}


//初始化应用
func (app *AppGaode) Init() {
	//初始化数据库
	app.Db = app.dbConnect()
	//校验数据表
	app.vailDbTables()
}


//数据库连接
func (app *AppGaode) dbConnect() (*gorm.DB) {
	conn := strings.Join([] string{
		app.DbConfig.Username , ":" ,
		app.DbConfig.Password , "@(" ,
		app.DbConfig.Hostname , ":" ,
		app.DbConfig.Port , ")/" ,
		app.DbConfig.Database , "?charset=" ,
		app.DbConfig.Charset , "&parseTime=True&loc=Local" ,
	} , "")
	//连接数据库
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}

	Log("db connect success .")

	db.LogMode(false)

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	//设置数据表名规则
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return app.DbConfig.Prefix + defaultTableName;
	}

	return db
}


//校验数据表
func (app *AppGaode) vailDbTables()  {
	if !app.Db.HasTable(&model.CongestionCitysRank{}) {
		if e := app.Db.CreateTable(&model.CongestionCitysRank{}).Error; e != nil {
			panic(e)
		}
	}

	Log("verification db tables .")
}


//运行任务
func (app *AppGaode) Run()  {

	if app.Stopdate == "" {
		Log("stopdate empty .")
		return
	}

	//解析开始时间
	starttime, e := datetime.DateStringToTime(app.Startdate , datetime.DATETIME)
	if e != nil {
		Log("starttime wrong format .");
		return
	}
	if starttime.Unix() <= 0 {
		Log("starttime wrong format .");
		return
	}

	//解析终止时间
	endtime, e := datetime.DateStringToTime(app.Stopdate , datetime.DATETIME)
	if e != nil {
		Log("stopdate wrong format .");
		return
	}
	if endtime.Unix() <= 0 {
		Log("stopdate wrong format .");
		return
	}

	//校验日期范围
	if starttime.Unix() >= endtime.Unix() {
		Log("date configuration error .");
		return
	}

	tips := false

	for {
		//当前日期
		nowtime := time.Now()
		//未到开始时间
		if starttime.Unix() > nowtime.Unix() {

			if tips == false {
				//提示等待中
				Log("waiting the task to begin .")
				tips = true;
			}

			continue;
		}
		//已到截止时间
		if endtime.Unix() <= nowtime.Unix() {
			Log("end of task .");
			break;
		}

		//执行获取
		Log("start gets ...")

		gaodeData := Gets(app.Ranks)

		Log("gets finish ...")

		//数据入库
		app.putStorage(gaodeData)

		Log("puts finish ...")

		Log("waiting " , app.Intervals , " seconds .")

		time.Sleep(time.Duration(app.Intervals) * time.Second) //等待间隔（s）
	}
}



//采集结果入库
func (app *AppGaode) putStorage(data [] *GaodeGetsData)  {
	lens := len(data)
	if lens == 0 {
		return
	}

	version := time.Now().Unix();

	for _ , gaode := range data {
		//创建文章列表
		rows := &model.CongestionCitysRank{
			City:gaode.City,
			Idx:gaode.Idx,
			RealSpeed:gaode.RealSpeed,
			FreeFlowSpeed:gaode.FreeFlowSpeed,
			Value:gaode.Value,
			Version:version,
		}

		if app.Db.NewRecord(rows) {
			app.Db.Create(&rows)
			Log("insert rows : " , rows.City , rows.ID)
		}
	}
}