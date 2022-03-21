package global

import (
	"myproject/pkg/logger"
	"myproject/pkg/setting"
)

var(
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting *setting.JWTSettings
	Logger *logger.Logger
	EmailSetting *setting.EmailSettings
)
