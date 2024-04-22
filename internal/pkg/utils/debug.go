package utils

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func PrintJson(v interface{}) {
	b, _ := jsoniter.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
