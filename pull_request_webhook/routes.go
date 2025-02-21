package infrastructure

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {

	routes := router.Group("pull_request")
	
	{
		routes.POST("/pull_resquest_event_process", func(ctx *gin.Context) {})
	}

	



	return
}