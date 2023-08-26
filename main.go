package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"blog-service/pkg/tracer"
	"fmt"
	"log"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func main() {
	fmt.Println("-----------" + global.ServerSetting.RunMode)
	r := routers.NewRouter()
	r.Run(":8080")

}

func setupSetting() error {
	//先使用setting套件裡的NewSetting將設定檔都讀入
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	//讀取區段設定，並存入指定載體
	err = setting.ReadSection("Server", &global.ServerSetting) //讀取區段Server，並存入global.ServerSetting結構體
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting) //讀取區段App，並存入global.AppSetting結構體
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting) //讀取區段Database，並存入global.DatabaseSetting結構體
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting) //讀取區段JWT，並存入global.JWTSetting結構體
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting) //讀取區段Email，並存入global.EmailSetting結構體

	//時間設定需加上時間單位
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	//啟用DB連線，並將指針傳遞到global.DBEngine變數
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 設定Logger
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{ //w io.Writer, prefix string, flag int
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
