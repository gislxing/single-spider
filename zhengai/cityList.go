package zhengai

import (
	"fmt"
	"regexp"
	"zbs/single-spider/engine"
)

var cityRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" data-v-0c63b635>([^<]+)</a>`)

func ParseCityList(contents []byte) (engine.ParseResult, error) {
	submatch := cityRe.FindAllSubmatch(contents, -1)
	if submatch == nil {
		return engine.ParseResult{}, fmt.Errorf("没有匹配的项")
	}

	parseResult := engine.ParseResult{}
	for _, v := range submatch {
		parseResult.Items = append(parseResult.Items, v[2])
		parseResult.Request = append(parseResult.Request, engine.Request{
			Url:       string(v[1]),
			ParseFunc: ParseCity,
		})
	}

	return parseResult, nil
}
