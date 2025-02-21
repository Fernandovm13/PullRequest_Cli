package infrastructure

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Handle_pull_request_event(ctx *gin.Context){

	eventType := ctx.GetHeader(X-GitHub-Event)
	deliveryD := ctx.GetHeader(X-GitHub-Delivery)

	log.Print()


}