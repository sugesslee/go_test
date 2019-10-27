package parser

import (
	"gostudy/crawler/engine"
	"gostudy/crawler/model"
	"regexp"
	"strconv"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/10/27     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/10/27 2:12 PM
 * @date 2019/10/27 2:12 PM
 * @since 1.0.0
 */
//var genderRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" .*>([\d]+)cm</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<)]+)</div>`)
var hukouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	//profile.Gender = extractString(contents, genderRe)
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err == nil {
		profile.Height = height
	}
	profile.Name = "昵称:" + name
	profile.Marriage = extractString(contents, marriageRe)
	//profile.Education = extractString(contents, educationRe)
	profile.Hukou = "工作地:" + extractString(contents, hukouRe)
	profile.Income = "月收入:" + extractString(contents, incomeRe)
	profile.Marriage = "婚况:" + extractString(contents, marriageRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
