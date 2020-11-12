package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lestrrat/go-file-rotatelogs"
	"log"
	"os"
	_ "startgin/docs"
	"startgin/routers"
	"time"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name youngxhu
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	godotenv.Load("conf/.env")

	// 设置分割日志
	//logWriterFormat := "var/log/gin.%Y%m%d.%H%M.log"
	// pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// fmt.Println(pwd,err)
	logWriterFormat := "var/log/gin.%Y%m%d.log"
	logWriter, err := rotatelogs.New(
		logWriterFormat,
		rotatelogs.WithLinkName("var/log/access_log.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}

	gin.DefaultWriter = logWriter
	log.Println(os.Getenv("DB_HOST"))
	engine := gin.Default()
	// engine.LoadHTMLGlob("resources/themes/***/***/*")
	// 初始化路由
	routers.InitRouter(engine)
	// 初始化模型
	// models.ConnectDB()
	engine.Run(":8080")
}
