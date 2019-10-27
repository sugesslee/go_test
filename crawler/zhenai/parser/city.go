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

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^>]+)</a>`

func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range match {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Request = append(result.Request, engine.Request{Url: string(m[1]), ParserFunc: func(c []byte) engine.ParserResult {
			return ParseProfile(c, name)
		},
		})
	}
	return result
}
