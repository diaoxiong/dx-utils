package regex

import "regexp"

func IsQuotedString(s string) bool {
	// 匹配单引号或双引号字符串的正则表达式
	quotedStringRegex := regexp.MustCompile(`^(?:"[^"]*"|'[^']*')$`)

	return quotedStringRegex.MatchString(s)
}

func IsInt64(s string) bool {
	r := regexp.MustCompile(`^-?\d{1,19}$`)
	return r.MatchString(s)
}
