package parser

import (
	"regexp"
	"spiders_on_go/crawler/engine"
	"spiders_on_go/crawler/model"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`.*album\.zhenai\.com/u/([\d]+)`)

// 函数接受两个参数，contents是用户资料页面的HTML内容的字节切片，name是用户的名字。函数返回一个engine.ParseResult类型的结果，包含解析出的用户资料。
// 主要作用是解析用户的个人资料页面，提取出用户的详细信息（如年龄、身高、月收入等），并将这些信息封装成一个model.Profile对象。然后，这个对象包装在engine.ParseResult中返回。
func ParseProfile(
	contents []byte, name string) engine.ParseResult {
	profile := model.Profile{} // 初始化一个model.Profile类型的变量profile，用于存储解析出的用户资料信息。
	profile.Name = name        // 将函数参数中传入的用户名赋值给profile的Name字段

	// 使用strconv.Atoi函数尝试将正则表达式匹配到的字符串转换为整数，并赋值给profile的相应字段（如年龄、身高、体重）。
	// 如果转换成功（err == nil），则更新profile对象的相应字段。
	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	// 对于那些不需要转换为整数的字段（如月收入、性别、是否购车等），直接调用extractString函数提取字符串，并赋值给profile的相应字段。
	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	// 初始化一个engine.ParseResult类型的变量result，并将profile对象添加到result.Items中。
	// 这样，解析出的用户资料就被封装在result中返回。
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	matches := guessRe.FindAllSubmatch(
		contents, -1)
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(
					c []byte) engine.ParseResult {
					return ParseProfile(c, name)
				},
			})
	}

	return result

}

// 该函数接受页面内容和一个正则表达式对象作为参数，用于从内容中提取匹配正则表达式的字符串。
// 如果找到匹配项，则返回第一个捕获组的内容；否则，返回空字符串。
func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}
