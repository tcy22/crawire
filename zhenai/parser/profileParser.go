package parser

import (
	"crawier/engine"
	"regexp"
	"crawier/model"
)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+岁)</td>`)
var marriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var professionRe  = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var genderRe  = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var hukouRe  = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parserProfile(contents []byte,url string,name string) engine.ParseResult{
	profile := model.Profile{}

	profile.Name = name
	profile.Age = extractString(contents,ageRe)
	profile.Gender = extractString(contents,genderRe)
	profile.Marriage = extractString(contents,marriageRe)
	profile.Education =extractString(contents,educationRe)
	profile.Hukou =extractString(contents,hukouRe)
	profile.Profession =extractString(contents,professionRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type:"zhenai",
				Id: extractString([]byte(url),idUrlRe),
				Payload:profile,
			},
		},
	}
	return result
}

func extractString(contents []byte,re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >=2 {
		return string(match[1])
	}else{
		return ""
	}
}