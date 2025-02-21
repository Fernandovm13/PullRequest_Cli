package main

import (
	infrastructure "weebhook_github/pull_request_webhook"

	"github.com/gin-gonic/gin"
)

func main() {

	Engine := gin.Default()

	infrastructure.Routes(Engine)
}