package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"myproject/global"
	"myproject/internal/model"
	"myproject/internal/routers"
	"myproject/pkg/logger"
	"myproject/pkg/setting"
	"net/http"
	"time"
)

func init()  {
	err := setupSetting()
	if err != nil{
		log.Fatalf("init.setupSetting err:%v",err)
	}
	err = setupLogger()
	if err != nil{
		log.Fatalf("init.setupLogger err:%v",err)
	}
	err = setupDBEngine()
	if err != nil{
		log.Fatalf("init.setupDBEngine err :%s",err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 项目
func main()  {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

//文件配置初始化
func setupSetting() error {
	s,err := setting.NewSetting()
	if err != nil{
		return err
	}
	err = s.ReadSection("Server",&global.ServerSetting)
	if err != nil{
		return err
	}
	err = s.ReadSection("App",&global.AppSetting)
	if err != nil{
		return err
	}
	err = s.ReadSection("Database",&global.DatabaseSetting)
	if err != nil{
		return err
	}
	err = s.ReadSection("jwt",&global.JWTSetting)
	if err != nil{
		return err
	}
	err = s.ReadSection("email",&global.EmailSetting)
	if err != nil{
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	return nil
}

//数据库初始化
func setupDBEngine() error {
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DatabaseSetting)
	if err  != nil{
		return err
	}
	return nil
}

//日志初始化
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavaPath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt,
		MaxSize: 60,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}
