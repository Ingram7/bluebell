package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1.获取参数
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 2.业务处理
	logic.SignUp(p)
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
