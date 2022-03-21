package app

import (
	"github.com/gin-gonic/gin"
	"myproject/pkg/errorcode"
	"net/http"
)

//响应处理
type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{
		Ctx: c,
	}
}

func (r *Response)ToResponse(data interface{})  {
	if data == nil{
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK,data)
}

func (r *Response)ToResponseList(list interface{},totalRows int)  {
	r.Ctx.JSON(http.StatusOK,gin.H{
		"list":list,
		"page":Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response)ToErrorResponse(err *errorcode.Error)  {
	response := gin.H{
		"code":err.Code(),
		"msg":err.Msg(),
	}
	details := err.Details()
	if len(details)>0{
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(),response)
}

