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

var (
	// punctuations 标点符号
	punctuations = []string{
		// 逗号
		"，", ",",
		// 句号
		"。", ".",
		// 感叹号
		"！", "!",
		// 问号
		"？", "?",
		// 冒号
		"：", ":",
		// 分号
		"；", ";",
		// 左/右单引号
		"‘", " '", "’", " '",
		// 左/右双引号
		"“", ` "`, "”", ` "`,
		// 左/右直角引号
		"「", " [", "」", " ]",
		"『", " [", "』", " ]",
		// 左/右括号
		"（", " (", "）", " )",
		"〔", " [", "〕", " ]",
		"【", " [", "】", " ]",
		"{", " {", "}", " }",
		// 省略号
		"……", "...",
		// 破折号
		"——", "-",
		// 连接号
		"—", "-",
		// 左/右斜杆
		"/", " /", "\\", " \\",
		// 波浪线
		"～", "~",
		// 书名号
		"《", " <", "》", " >",
		"〈", " <", "〉", " >",
		// 间隔号
		"·", " ·",
		// 顿号
		"、", ",",
	}
	// finals 韵母表
	finals = []string{
		// a
		"a1", "ā", "a2", "á", "a3", "ǎ", "a4", "à",
		// o
		"o1", "ō", "o2", "ó", "o3", "ǒ", "o4", "ò",
		// e
		"e1", "ē", "e2", "é", "e3", "ě", "e4", "è",
		// i
		"i1", "ī", "i2", "í", "i3", "ǐ", "i4", "ì",
		// u
		"u1", "ū", "u2", "ú", "u3", "ǔ", "u4", "ù",
		// v
		"v1", "ǖ", "v2", "ǘ", "v3", "ǚ", "v4", "ǜ",

		// ai
		"ai1", "āi", "ai2", "ái", "ai3", "ǎi", "ai4", "ài",
		// ei
		"ei1", "ēi", "ei2", "éi", "ei3", "ěi", "ei4", "èi",
		// ui
		"ui1", "uī", "ui2", "uí", "ui3", "uǐ", "ui4", "uì",
		// ao
		"ao1", "āo", "ao2", "áo", "ao3", "ǎo", "ao4", "ào",
		// ou
		"ou1", "ōu", "ou2", "óu", "ou3", "ǒu", "ou4", "òu",
		// iu
		"iu1", "īu", "iu2", "íu", "iu3", "ǐu", "iu4", "ìu",

		// ie
		"ie1", "iē", "ie2", "ié", "ie3", "iě", "ie4", "iè",
		// ve
		"ue1", "üē", "ue2", "üé", "ue3", "üě", "ue4", "üè",
		// er
		"er1", "ēr", "er2", "ér", "er3", "ěr", "er4", "èr",

		// an
		"an1", "ān", "an2", "án", "an3", "ǎn", "an4", "àn",
		// en
		"en1", "ēn", "en2", "én", "en3", "ěn", "en4", "èn",
		// in
		"in1", "īn", "in2", "ín", "in3", "ǐn", "in4", "ìn",
		// un/vn
		"un1", "ūn", "un2", "ún", "un3", "ǔn", "un4", "ùn",

		// ang
		"ang1", "āng", "ang2", "áng", "ang3", "ǎng", "ang4", "àng",
		// eng
		"eng1", "ēng", "eng2", "éng", "eng3", "ěng", "eng4", "èng",
		// ing
		"ing1", "īng", "ing2", "íng", "ing3", "ǐng", "ing4", "ìng",
		// ong
		"ong1", "ōng", "ong2", "óng", "ong3", "ǒng", "ong4", "òng",
	}
)

// -----------------------------------------------------------------------------

//转换结果
type ConvertResult string

//创建转换结果对象
func NewConvertResult(s string) *ConvertResult {
	cr := ConvertResult(s)
	return &cr
}

// ASCII 带数字的声调
func (r *ConvertResult) ASCII() string {
	return string(*r)
}

// Unicode Unicode声调
func (r *ConvertResult) Unicode() string {
	s := string(*r)
	for i := len(finals) - 1; i >= 0; i -= 2 {
		s = strings.Replace(s, finals[i-1], finals[i], -1)
	}
	return s
}

// None 不带声调输出
func (r *ConvertResult) None() string {
	s := string(*r)
	re := regexp.MustCompile(`[1-4]{1}`)
	s = re.ReplaceAllString(s, "")
	return s
}
