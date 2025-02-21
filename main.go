package main

import (
	"log"

    infrastructure "weebhook_github/pull_request_webhook"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	infrastructure.Routes(router)

	log.Println("Servidor corriendo en el puerto 8080")
	router.Run(":8080")
}
