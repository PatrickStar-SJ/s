package parser

import (
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	/*
		// 先执行下面这段，将结果复制粘贴落盘成   citylist_test_data.html文件
		contents, err := fetcher.Fetch(
			"http://localhost:8080/mock/www.zhenai.com/zhenghun")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", contents)
	*/

	contents, err := os.ReadFile(
		"citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)

	result := ParseCityList(contents)
	//fmt.Println(len(result.Requests))  //437

	const resultSize = 437
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/binhaixin",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/binzhou",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/bishan",
	}
	expectedCities := []string{
		"City 滨海新", "City 滨州", "City 璧山",
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Items))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but"+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but "+
				"was %s",
				i, city, result.Items[i].(string))
		}
	}

}
