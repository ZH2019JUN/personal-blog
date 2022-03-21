package util

import (
	"crypto/md5"
	"encoding/hex"
)

//上传文件的工具库


//对上传的文件名进行格式化
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}


