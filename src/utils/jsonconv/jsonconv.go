package jsonconv

import "encoding/json"

// 包内方法
func ObjectToJson(data interface{}) string {
	bytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

type Jsonable interface {
	ToJson() string
}
