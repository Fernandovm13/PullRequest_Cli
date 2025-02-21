package application

import (
	"encoding/json"
	"log"
    "weebhook_github/pull_request_webhook/domain/value_objects"
)

func ProcessPullRequestEvent(payload []byte) {
	var prPayload valueobjects.PullRequestPayload

	err := json.Unmarshal(payload, &prPayload)
	if err != nil {
		log.Printf("Error al parsear el JSON: %v", err)
		return
	}

	if prPayload.Action == "closed" {
		log.Printf("Pull Request cerrado:\n - Desde: %s\n - Hacia: %s\n - Usuario: %s\n - Repositorio: %s\n - URL: %s",
			prPayload.PullRequest.Head.Ref,
			prPayload.PullRequest.Base.Ref,
			prPayload.PullRequest.User.Login,
			prPayload.PullRequest.Repository.Name,
			prPayload.PullRequest.URL,
		)
	}
}
