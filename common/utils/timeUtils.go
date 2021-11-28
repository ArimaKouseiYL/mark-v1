package utils

import (
	"strconv"
	"time"
)

func GetNowDays() string {

	now := time.Now()                  //获取当前时间
	timestamp := now.Unix()            //时间戳
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	year := timeObj.Year()             //年
	month := timeObj.Month()           //月
	day := timeObj.Day()               //日

	return strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
}
