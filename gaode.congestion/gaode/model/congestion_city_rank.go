package model

import (
	"time"
)

//拥堵城市排行榜
type CongestionCitysRank struct {
	ID  uint `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"type:datetime;"`
	UpdatedAt time.Time `gorm:"type:datetime;"`
	DeletedAt *time.Time `gorm:"type:datetime"`

	City string `gorm:"type:varchar(25);default:''"` //城市名称
	Idx string `gorm:"type:char(8);"` //拥堵指数
	RealSpeed string `gorm:"type:char(8);"`  //平均行车速度（km/h）
	FreeFlowSpeed string `gorm:"type:char(8);"`  //畅通行车速度（km/h）
	Value int `gorm:"type:int(8);"`  //城市排名
	Version int64 `gorm:"type:int(11);default:0"` //入库版本
}