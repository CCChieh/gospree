package util

import (
	"github.com/json-iterator/go"
	"time"
)

//获取json中日志的时间
func GetJsonTime(js []byte) time.Time {
	//匿名结构体，存储json中的时间
	v := &struct {
		Time time.Time
	}{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(js, v)
	if err != nil {
		panic(err.Error())
	}
	return v.Time
}
