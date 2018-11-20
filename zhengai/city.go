package zhengai

import (
	"fmt"
	"regexp"
	"strconv"
	"zbs/single-spider/engine"
	"zbs/single-spider/model"
)

var userListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var sexRe = regexp.MustCompile(`<td width="[0-9]+"><span class="grayL">性别：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="grayL">居住地：</span>([^<]+)</td>`)
var ageRe = regexp.MustCompile(`<td width="[0-9]+"><span class="grayL">年龄：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td width="[0-9]+"><span class="grayL">婚况：</span>([^<]+)</td>`)

func ParseCity(contents []byte) (engine.ParseResult, error) {
	submatch := userListRe.FindAllSubmatch(contents, -1)
	if submatch == nil {
		return engine.ParseResult{}, fmt.Errorf("no match item")
	}

	sexMatch := sexRe.FindAllSubmatch(contents, -1)
	houseMatch := houseRe.FindAllSubmatch(contents, -1)
	ageMatch := ageRe.FindAllSubmatch(contents, -1)
	marriageMatch := marriageRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for i, name := range submatch {
		age, _ := strconv.Atoi(string(ageMatch[i][1]))

		result.Items = append(result.Items, model.User{
			Name:     string(name[2]),
			Sex:      string(sexMatch[i][1]),
			House:    string(houseMatch[i][1]),
			Age:      age,
			Marriage: string(marriageMatch[i][1]),
		})
	}

	return result, nil
}
