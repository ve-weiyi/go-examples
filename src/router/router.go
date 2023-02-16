package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/go-examplse/src/api"
	"github.com/ve-weiyi/go-examplse/src/utils/response"
	"log"
)

// InitRouter router 入口文件
func InitRouter() *gin.Engine {
	log.Printf("开始注册api路径......")

	engine := gin.New()
	group := engine.Group("/api/v1")
	//处理路由
	group.GET("/hello", func(ctx *gin.Context) {
		//以字符串格式返回
		ctx.JSON(200, "hello world!")
	})

	group.GET("/login", func(ctx *gin.Context) {
		name := ctx.GetString("username")
		pwd := ctx.GetString("password")

		login, err := api.Login(name, pwd)
		if err != nil {
			ctx.JSON(200, response.ErrorResponse{
				Response: response.Response{
					Code:    0,
					Message: "",
				},
				Data: err,
			})
			return
		}
		//以字符串格式返回
		ctx.JSON(200, response.SuccessResponse{
			Response: response.Response{
				Code:    501,
				Message: "success",
			},
			Data: login,
		})
	})

	//group.POST("/register", api.Register)

	return engine
}
