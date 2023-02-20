package response

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
