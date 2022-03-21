package service

import (
	"errors"
	"mime/multipart"
	"myproject/global"
	"myproject/pkg/upload"
	"os"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

func (svc *Service)UploadFile(fileType upload.FileType,file multipart.File,fileHeader *multipart.FileHeader) (*FileInfo,error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath+"/"+fileName
	if upload.CheckContainExt(fileType,fileName){
		return nil,errors.New("file suffix is not supported")
	}
	if upload.CheckSavePath(uploadSavePath){
		err := upload.CreateSavePath(uploadSavePath,os.ModePerm)
		if err != nil{
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckMAxSize(fileType,file){
		return nil,errors.New("exceeded maximum file limit")
	}
	if err := upload.SaveFile(fileHeader,dst);err != nil{
		return nil, err
	}
	if upload.CheckPermission(uploadSavePath){
		return nil,errors.New("insufficient file permission")
	}
	accessUrl := global.AppSetting.UploadServerUrl+"/"+fileName
	return &FileInfo{Name: fileName,AccessUrl: accessUrl},nil
}
