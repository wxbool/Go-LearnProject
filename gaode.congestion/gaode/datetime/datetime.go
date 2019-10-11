package datetime

import "time"

const (
	DATE = 1 //字符串日期类型：2008-06-01
	DATETIME = 2 //字符串日期类型：2008-06-01 12:01:01
)


func DateStringToTime (value string , valueType int) (time.Time , error) {
	timeLayout := ""
	switch valueType {
		case DATE:
			timeLayout = "2006-01-02"
		case DATETIME:
			timeLayout = "2006-01-02 15:04:05"
	}

	loc , e := time.LoadLocation("Local") //本地时区
	if e != nil {
		return time.Now() , e
	}
	time , e := time.ParseInLocation(timeLayout , value , loc)
	if e != nil {
		return time , e
	}
	
	return time , nil
}