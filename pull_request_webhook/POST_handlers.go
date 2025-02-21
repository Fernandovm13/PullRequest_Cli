package infrastructure

import (
	"io"
	"net/http"
	"weebhook_github/pull_request_webhook/application"

	"github.com/gin-gonic/gin"
)

func HandlePullRequestEvent(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")

	if eventType != "pull_request" {
		ctx.JSON(http.StatusOK, gin.H{"message": "Evento ignorado"})
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el cuerpo de la petición"})
		return
	}

	application.ProcessPullRequestEvent(body)
	ctx.JSON(http.StatusOK, gin.H{"message": "Evento procesado correctamente"})
}
