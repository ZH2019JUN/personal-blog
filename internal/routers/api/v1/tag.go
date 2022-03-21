package v1

import (
	"github.com/gin-gonic/gin"
	"myproject/global"
	"myproject/internal/service"
	"myproject/pkg/app"
	"myproject/pkg/convert"
	"myproject/pkg/errorcode"
)

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	//参数绑定
	valid,errs := app.BindingAndValid(c,&param)
	if !valid{
		global.Logger.Errorf(c,"app.BindingAndValid errs: %v",errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c),PageSize: app.GetPageSize(c)}
	totalRows,err := svc.CountTag(&service.CountTagRequest{Name: param.Name,State: param.State})
	if err != nil{
		global.Logger.Errorf(c,"svc.CountTag err: %v",err)
		response.ToErrorResponse(errorcode.ErrorCountTagFail)
		return
	}

	tags,err := svc.GetTagList(&param,&pager)
	if err != nil{
		global.Logger.Errorf(c,"svc.GetTagList err: %v",err)
		response.ToErrorResponse(errorcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags,totalRows)
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindingAndValid(c,&param)
	if !valid{
		global.Logger.Errorf(c,"app.BindingAndValid err: %v",errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil{
		global.Logger.Errorf(c,"CreateTag err: %v",err)
		response.ToErrorResponse(errorcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid,errs := app.BindingAndValid(c,&param)
	if !valid{
		global.Logger.Errorf(c,"app.BindingAndValid err: %v",errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil{
		global.Logger.Errorf(c,"UpdateTag err: %v",err)
		response.ToErrorResponse(errorcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid,errs := app.BindingAndValid(c,&param)
	if !valid{
		global.Logger.Errorf(c,"app.BindingAndValid err: %v",errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil{
		global.Logger.Errorf(c,"DeleteTag err: %v",err)
		response.ToErrorResponse(errorcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
