package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["China"] = "北京"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 查看元素是否在集合中
	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 首都是：", capital)
	} else {
		fmt.Println("American 首都不存在")
	}

	delete(countryCapitalMap, "France")
	fmt.Println(countryCapitalMap)
}
