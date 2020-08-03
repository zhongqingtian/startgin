package routers

import (
	"github.com/gin-gonic/gin"
	"startgin/routers/api/v1/demo"
)

// ApiV1 接口v1版本路由
func ApiV1(engine *gin.Engine) {
	g := engine.Group("/api/v1")
	// 简单 get post put delete请求示例
	g.GET("/demo/simple/get", demo.Get)
	g.POST("/demo/simple/post", demo.Post)
	g.PUT("/demo/simple/put", demo.Put)
	g.DELETE("/demo/simple/delete", demo.Delete)
}

// ApiV2 接口v2版本路由
func ApiV2(engine *gin.Engine) {
	/*g := engine.Group("/api/v2")
	{
		g.GET("/jenkins/create",jenkins.CreateJob) // 创建一个job
		g.GET("/jenkins/delete",jenkins.DeleteJob)
		g.GET("/jenkins/build",jenkins.BuildJob) // 构建job
		g.GET("/jenkins/getjob",jenkins.GetJob) // 获取job
		g.GET("/jenkins/search",jenkins.SearchJobs) // 查找
		g.GET("/jenkins/category",jenkins.GetJobType) // 可构建类型列表
		g.GET("/jenkins/history",jenkins.GetHistory) // 获取某个job的历史构建信息
		g.GET("jenkins/log",jenkins.Log) // 获取某个job的构建日志
	}*/
}

// ApiV3 接口v3版本路由
func ApiV3(engine *gin.Engine) {
	//g := engine.Group("/api/v3")
	//{
	//	g.GET("/demo/json")
	//}
}