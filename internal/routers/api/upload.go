package api

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/service"
	"myproject/pkg/app"
	"myproject/pkg/convert"
	"myproject/pkg/errorcode"
	"myproject/pkg/upload"
)

type Upload struct {}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload)UploadFile(c *gin.Context)  {
	response := app.NewResponse(c)
	file,fileHeader,err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustUInt32()
	if err != nil{
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0{
		response.ToErrorResponse(errorcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo,err := svc.UploadFile(upload.FileType(fileType),file,fileHeader)
	if err != nil{
		response.ToErrorResponse(errorcode.ERROR_UPLOAD_FILE_FAIL)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url":fileInfo.AccessUrl,
	})
}
