//++++++++++++++++++++++++++++++++++++++++
// https://gitee.com/shirdonl
//++++++++++++++++++++++++++++++++++++++++
// Author:ShirDon
//++++++++++++++++++++++++++++++++++++++++

package library

import (
	"regexp"
	"strings"
)

//拼音词典
type Dictionary struct{}

//新建拼音词典对象
func NewDict() *Dictionary {
	return &Dictionary{}
}

//中文转换为拼音, 不保留标点符号
func (p *Dictionary) Convert(s string, sep string) (result *ConvertResult) {
	s = p.toRoman(s, false)

	split := ToSlice(s)

	result = NewConvertResult(strings.Join(split, sep))
	return
}

//中文转换为拼音, 保留标点符号
func (p *Dictionary) Sentence(s string) (result *ConvertResult) {
	s = p.toRoman(s, false)

	r := regexp.QuoteMeta(strings.Join(punctuations, ""))
	r = strings.Replace(r, " ", "", -1)
	re := regexp.MustCompile("[^a-zA-Z0-9" + r + `\s_]+`)
	s = re.ReplaceAllString(s, "")

	for i := 0; i < len(punctuations); i += 2 {
		s = strings.Replace(s, punctuations[i], punctuations[i+1], -1)
	}

	result = NewConvertResult(s)
	return
}

//转换人名
func (p *Dictionary) Name(s string, sep string) (result *ConvertResult) {
	s = p.toRoman(s, true)

	split := ToSlice(s)

	result = NewConvertResult(strings.Join(split, sep))
	return
}

//获取拼音的首字符
func (p *Dictionary) Abbr(s string, sep string) string {
	s = p.toRoman(s, false)

	var abbr []string
	for _, item := range ToSlice(s) {
		abbr = append(abbr, item[0:1])
	}

	return strings.Join(abbr, sep)
}

func (p *Dictionary) prepare(s string) string {
	var re *regexp.Regexp

	re = regexp.MustCompile(`[a-zA-Z0-9_-]+`)
	s = re.ReplaceAllStringFunc(s, func(repl string) string {
		return "\t" + repl
	})

	re = regexp.MustCompile(`[^\p{Han}\p{P}\p{Z}\p{M}\p{N}\p{L}\t]`)
	s = re.ReplaceAllString(s, "")

	return s
}

func (p *Dictionary) toRoman(s string, convertName bool) string {
	s = p.prepare(s)

	if convertName {
		for i := 0; i < len(surnames); i += 2 {
			if strings.Index(s, surnames[i]) == 0 {
				s = strings.Replace(s, surnames[i], surnames[i+1], 1)
			}
		}
	}

	for i := 0; i < len(dict); i += 2 {
		s = strings.Replace(s, dict[i], dict[i+1], -1)
	}

	s = strings.Replace(s, "\t", " ", -1)
	s = strings.Replace(s, "  ", " ", -1)
	s = strings.TrimSpace(s)

	return s
}

//转换为字符串数组
func ToSlice(s string) []string {
	var split []string
	re := regexp.MustCompile(`[^a-zA-Z1-4]+`)
	for _, str := range re.Split(s, -1) {
		if str != "" {
			split = append(split, str)
		}
	}
	return split
}
