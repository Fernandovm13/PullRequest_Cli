package valueobjects

type PullRequestPayload struct {
	Action      string `json:"action"`
	PullRequest struct {
		URL     string `json:"html_url"`
		Number  int    `json:"number"`
		Head    struct {
			Ref string `json:"ref"`
		} `json:"head"`
		Base struct {
			Ref string `json:"ref"`
		} `json:"base"`
		User struct {
			Login string `json:"login"`
		} `json:"user"`
		Repository struct {
			Name string `json:"name"`
		} `json:"repository"`
	} `json:"pull_request"`
}
