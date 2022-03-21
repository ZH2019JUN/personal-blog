package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize int
	LogSavaPath string
	LogFileName string
	LogFileExt string
	//上传文件的相关属性
	UploadSavePath string
	UploadServerUrl string
	UploadImageSize int
	//上传文件所允许的文件后缀
	UploadImageAllowExts []string
	DefaultContextTimeout time.Duration
}

type DatabaseSettingS struct {
	DBType string
	Username string
	Password string
	Host string
	DBName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettings struct {
	Host string
	Port int
	IsSSL bool
	UserName string
	Password string
	From string
	To []string
}

var sections = make(map[string]interface{})

func (s *Setting)ReadSection(k string,v interface{}) error {
	err := s.vp.UnmarshalKey(k,v)
	if err != nil{
		return err
	}
	if _,ok := sections[k];!ok{
		sections[k] = v
	}
	return nil
}

func (s *Setting)ReloadAllSection() error {
	for k,v := range sections{
		err := s.ReadSection(k,v)
		if err != nil{
			return err
		}
	}
	return nil
}
