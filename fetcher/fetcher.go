package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
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

// 拉取网页源码
// 部分网站有反爬虫，所以此处需要设置请求头中的 User-Agent
func FetchUseHeader(url string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(request)
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
