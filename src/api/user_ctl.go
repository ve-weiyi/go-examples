package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/go-examplse/src/model"
	"github.com/ve-weiyi/go-examplse/src/utils/response"
)

func Login(name, pwd string) (*model.User, error) {
	db := database.DB()
	var user model.User
	user.Username = name
	user.Password = pwd
	//验证格式
	if !user.IsValid() {
		return nil, errors.New("账号或密码格式错误")
	}
	//查询数据库
	err := db.Where("username = ? and password = ?", name, pwd).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Register(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(200, response.NewFailResponse(err))
		return
	}

}
