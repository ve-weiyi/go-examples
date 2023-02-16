package response

import (
	"encoding/json"
	"errors"
)

type Response struct {
	Code    int
	Message string
}

type SuccessResponse struct {
	Response //继承
	//Resp     Response //组合
	Data interface{}
}

type ErrorResponse struct {
	Response //继承
	//Response Response //组合
	Data error
}

// 结构体方法
func (m *SuccessResponse) ToJson() string {
	bytes, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (m *SuccessResponse) GetAddress() (*Response, error) {
	if m.Code == 200 {
		return nil, errors.New("返回错误")
	}
	return &m.Response, nil
}

func NewOkResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Response: Response{
			Code:    200,
			Message: "success",
		},
		Data: data,
	}
}

func NewFailResponse(data interface{}) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Code:    200,
			Message: "success",
		},
		Data: data.(error),
	}
}
