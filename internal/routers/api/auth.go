package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"myproject/global"
	"myproject/internal/service"
	"myproject/pkg/app"
	"myproject/pkg/errorcode"
)

func GetAuth(c *gin.Context)  {
	param := service.AuthRequest{}
	response:= app.NewResponse(c)
	valid,errs := app.BindingAndValid(c,&param)
	if !valid {
		global.Logger.Errorf(c,"app.BindingAndValid errs :%v",errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil{
		global.Logger.Errorf(c,"svc.CheckAuth err :%v",err)
		response.ToErrorResponse(errorcode.UnauthorizedAuthNotExist)
		return
	}
	token,err := app.GenerateToken(param.AppKey,param.AppSecret)
	if err != nil{
		log.Fatal(err )
		global.Logger.Errorf(c,"app.GenerateToken err :%v",err)
		response.ToErrorResponse(errorcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token":token,
	})
	return
}
