package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

// 拉取网页源码
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("网址: %v 响应码错误: %v", url, resp.StatusCode)
	}

	// 获取网页编码格式
	bufReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bufReader)
	reader := transform.NewReader(bufReader, e.NewDecoder())
	return ioutil.ReadAll(reader)
}

// 猜测页面的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
