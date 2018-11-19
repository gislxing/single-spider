package zhengai

import (
	"fmt"
	"regexp"
	"zbs/single-spider/engine"
)

var userListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var xx = regexp.MustCompile(`<td width="[0-9]+"><span class="grayL">性别：</span>([^<]+)</td>`)

func ParseCity(contents []byte) (engine.ParseResult, error) {
	submatch := userListRe.FindAllSubmatch(contents, -1)
	if submatch == nil {
		return engine.ParseResult{}, fmt.Errorf("no match item")
	}

	result := engine.ParseResult{}
	for _, value := range submatch {
		result.Items = append(result.Items, value[2])
	}

	return result, nil
}
