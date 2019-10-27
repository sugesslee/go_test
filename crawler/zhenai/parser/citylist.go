package parser

import (
	"gostudy/crawler/engine"
	"regexp"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/10/27     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/10/27 10:43 AM
 * @date 2019/10/27 10:43 AM
 * @since 1.0.0
 */

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	limit := 10

	for _, m := range match {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Request = append(result.Request, engine.Request{Url: string(m[1]), ParserFunc: ParseCity})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
