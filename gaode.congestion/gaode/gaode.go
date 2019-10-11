package gaode

import (
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"strconv"
)

type GaodeGetsData struct {
	City string  //城市名称
	Idx string  //拥堵指数
	RealSpeed string  //平均行车速度（km/h）
	FreeFlowSpeed string //畅通行车速度（km/h）
	Value int //城市排名
}


func CheckError (e error) {
	if e != nil {
		panic(e)
	}
}


//执行一次获取
func Gets(ranks int) ([] *GaodeGetsData) {

	targetApi := "https://report.amap.com/ajax/getCityRank.do"

	request, e := http.NewRequest(http.MethodGet, targetApi , nil)
	CheckError(e)

	request.Header.Set("Host" , "report.amap.com")
	request.Header.Set("Referer" , "https://report.amap.com/index.do")
	request.Header.Set("User-Agent" , "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.204 Safari/537.36")

	response, e := http.DefaultClient.Do(request)
	CheckError(e)

	content, e := ioutil.ReadAll(response.Body)
	CheckError(e)

	if string(content) == "" {
		panic("gets error .")
	}

	//声明结果队列
	result := make([] *GaodeGetsData , 0)
	index := 0 //获取计数
	//遍历
	jsonparser.ArrayEach(content , func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		index++

		if index > ranks {
			return //跳过获取
		}

		gaode := &GaodeGetsData{}

		city , _ := jsonparser.GetString(value, "label")
		idx , _ := jsonparser.GetFloat(value, "idx")
		realSpeed , _ := jsonparser.GetFloat(value, "realSpeed")
		freeFlowSpeed , _ := jsonparser.GetFloat(value, "freeFlowSpeed")
		valueint , _ := jsonparser.GetInt(value, "value")

		gaode.City = city
		gaode.Idx  = strconv.FormatFloat(idx, 'f', 2, 64)
		gaode.RealSpeed  = strconv.FormatFloat(realSpeed, 'f', 2, 64)
		gaode.FreeFlowSpeed  = strconv.FormatFloat(freeFlowSpeed, 'f', 2, 64)
		gaode.Value  = int(valueint)

		result = append(result , gaode)
	})

	return result
}

