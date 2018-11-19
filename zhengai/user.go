package zhengai

import (
	"regexp"
	"zbs/single-spider/engine"
)

var userRe = regexp.MustCompile(`<div data-v-35c72236="" class="des f-cl">阿坝 | 41岁 | 大学本科 | 离异 | 163cm | 5001-8000元<a data-v-35c72236="" href="https://www.zhenai.com/n/login?fromurl=http%3A%2F%2Falbum.zhenai.com%2Fu%2F1572218980" target="_self" class="online f-fr">查看最后登录时间</a></div>`)

func ParseUser(contents []byte) (engine.ParseResult, error) {
	//user := model.User{}
	return engine.ParseResult{}, nil
}
