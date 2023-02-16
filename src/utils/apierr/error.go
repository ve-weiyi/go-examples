package apierr

import (
	"encoding/json"
	"errors"
)

type Url struct {
	Host string
	Port int
}

type ApiError struct {
	Url //继承
	//UrlAddr Url //组合
	Code int
	Msg  string
	Data interface{}
}

// 结构体方法
func (m *ApiError) ToJson() string {
	bytes, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (m *ApiError) GetAddress() (*Url, error) {
	if m.Code == 200 {
		return nil, errors.New("返回错误")
	}
	return &m.Url, nil
}
