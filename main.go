package main

import (
	"zbs/single-spider/engine"
	"zbs/single-spider/zhengai"
)

const startUrl = `http://www.zhenai.com/zhenghun`

func main() {
	engine.Run(engine.Request{
		Url:       startUrl,
		ParseFunc: zhengai.ParseCityList,
	})
}
