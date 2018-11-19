package zhengai

import (
	"fmt"
	"regexp"
	"zbs/single-spider/engine"
)

var userListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte) (engine.ParseResult, error) {
	submatch := userListRe.FindAllSubmatch(contents, -1)
	if submatch == nil {
		return engine.ParseResult{}, fmt.Errorf("no match item")
	}

	result := engine.ParseResult{}
	for _, value := range submatch {
		result.Items = append(result.Items, value[2])
		result.Request = append(result.Request, engine.Request{
			Url: string(value[2]),
			ParseFunc: func(bytes []byte) (engine.ParseResult, error) {
				return engine.ParseResult{}, nil
			},
		})
	}

	return result, nil
}
