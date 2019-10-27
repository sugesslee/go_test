package parser

import (
	"io/ioutil"
	"testing"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/10/27     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/10/27 1:30 PM
 * @date 2019/10/27 1:30 PM
 * @since 1.0.0
 */
func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	result := ParseCityList(contents)

	// verify result
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(result.Request) != resultSize {
		t.Errorf("result should have %d requests; but get %d request.", resultSize, len(result.Request))
	}

	for i, url := range expectedUrls {
		if result.Request[i].Url != url {
			t.Errorf("expected url %d: %s; but was %s", i, url, result.Request[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items; but get %d items.", resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city %d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}
}
