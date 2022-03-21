package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"myproject/global"
	"myproject/pkg/util"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

//获取加密的文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name,ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

//获取路径使用的文件扩展名
func GetFileExt(name string) string {
	return path.Ext(name)
}

//获取文件保存地址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//检查文件的相关方法

//检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_,err := os.Stat(dst)
	return os.IsNotExist(err)
}

//检查文件后缀是否包含在约定的后缀之中
func CheckContainExt(t FileType,name string) bool {
	ext := GetFileExt(name)
	switch t{
	case TypeImage:
		for _,allowExt := range global.AppSetting.UploadImageAllowExts{
			if strings.ToUpper(allowExt) == strings.ToUpper(ext){
				return true
			}
		}
	}
	return false
}

//检查文件大小是否超出限制
func CheckMAxSize(t FileType,f multipart.File) bool {
	content,_ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageSize{
			return true
		}
	}
	return false
}

//检查文件权限
func CheckPermission(dst string) bool {
	_,err := os.Stat(dst)
	return os.IsPermission(err)
}

//对文件进行写入和创建等相关操作

//创建保存上传文件的目录
func CreateSavePath(dst string,perm os.FileMode) error {
	err := os.Mkdir(dst,perm)
	if err != nil{
		return err
	}
	return nil
}

//保存上传的文件
func SaveFile(file *multipart.FileHeader,dst string) error {
	src,err := file.Open()
	if err != nil{
		return err
	}
	defer src.Close()
	out,err := os.Create(dst)
	if err != nil{
		return err
	}
	defer out.Close()
	_,err = io.Copy(out,src)
	return err
}

