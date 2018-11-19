package zhengai

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile(`cityList_test_data.html`)
	if err != nil {
		panic(err)
	}

	result, err := ParseCityList(contents)
	if err != nil {
		panic(err)
	}

	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d, but had %d\n", resultSize, len(result.Items))
	}

	for i, url := range expectedUrls {
		if result.Request[i].Url != url {
			t.Errorf("expected url #%d: %s, but %s", i, url, result.Request[i].Url)
		}
	}

	if len(result.Request) != resultSize {
		t.Errorf("result should have %d, bud had %d\n", resultSize, len(result.Request))
	}

	for i, city := range expectedCities {
		if fmt.Sprintf("%s", result.Items[i]) != city {
			t.Errorf("expected ciyt #%d: %v, but %v", i, city, fmt.Sprintf("%s", result.Items[i]))
		}
	}
}
