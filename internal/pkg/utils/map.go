package utils

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func GetMapValue(data map[string]interface{}, key string) (interface{}, bool) {
	if data == nil {
		return nil, false
	}
	keys := strings.Split(key, ".")
	var value interface{} = data

	for _, k := range keys {
		switch typedValue := value.(type) {
		case map[string]interface{}:
			nextValue, ok := typedValue[k]
			if !ok {
				return nil, false
			}
			value = nextValue
		default:
			return nil, false
		}
	}

	return value, true
}

func ToMap(data interface{}) (map[string]interface{}, error) {

	var (
		jsonRaw []byte
		e       error
	)
	if raw, ok := data.([]byte); ok {
		jsonRaw = raw
	} else {
		jsonRaw, e = jsoniter.Marshal(data)
		if e != nil {
			return nil, e
		}
	}

	var res map[string]interface{}
	e = jsoniter.Unmarshal(jsonRaw, &res)
	return res, e
}
