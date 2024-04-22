package utils

import (
	"fmt"
	"regexp"
)

func FillString(s string, lenght int, placeholder ...string) string {
	if len(s) >= lenght {
		return s
	}
	if len(placeholder) == 0 {
		placeholder = append(placeholder, " ")
	}
	for len(s) < lenght {
		s += placeholder[0]
	}
	return s
}

func UnwrapBytesString(value []byte) []byte {
	l := len(value)
	if l > 1 {
		if value[0] == '"' && value[l-1] == '"' {
			if l == 2 {
				return nil
			}
			return value[1 : l-2]
		}
	}

	return value
}

func LimitString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

// 解析模板字符串
func ParseTemplate(template string, obj map[string]interface{}) string {
	re := regexp.MustCompile(`{(\S+?)}`) // 匹配形式为{var}的字符串
	result := re.ReplaceAllStringFunc(template, func(match string) string {
		prop := re.FindStringSubmatch(match)[1] // 提取匹配的属性名

		if value, ok := GetMapValue(obj, prop); ok {
			switch vt := value.(type) {
			case string:
				return vt
			case int, int32, int64:
				return fmt.Sprintf("%d", vt)
			case float64, float32:
				return fmt.Sprintf("%f", vt)
			default:
				return match // 如果属性值不是字符串、整数或浮点数，则保持原样
			}
		}
		return match // 如果未找到对应属性，则保持原样
	})

	return result
}
