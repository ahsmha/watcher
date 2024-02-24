package github

type GithubHookRequest struct {
	Ref     string `json:"ref"`
	Before  string `json:"before"`
	After   string `json:"after"`
	Commits []struct {
		Id      string `json:"id"`
		Message string `json:"message"`
		Author  struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
	} `json:"commits"`
}
