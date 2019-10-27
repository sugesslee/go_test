package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com@abc.com
email1 is aaa@123.com.cn
email2 is 345@qq.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	match := re.FindAllStringSubmatch(text, -1)

	for _, m := range match {
		fmt.Println(m)
	}
}
