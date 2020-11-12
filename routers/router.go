package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
func InitRouter(engine *gin.Engine) *gin.Engine {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// api 接口路由配置
	ApiV1(engine)
	ApiV2(engine)
	ApiV3(engine)

	// app 应用路由配置
	/*AppV1(engine)
	AppV2(engine)
	AppV3(engine)*/
	return engine
}
