package github

import "time"

type PushEventRequest struct {
	Source  string
	Type    string
	Commits []PushCommit `json:"commits"`
	Payload PushEventRequestPayload
}

type PushEventRequestPayload struct {
	Repository interface{} `json:"repository"`
	Pusher     interface{} `json:"pusher"`
	Sender     interface{} `json:"sender"`
	Head       interface{} `json:"head_commit"`
}

type PushCommit struct {
	ID        string        `json:"id"`
	TreeID    string        `json:"tree_id"`
	URL       string        `json:"url"`
	Committer PushCommitter `json:"author"`
	Timestamp *time.Time    `json:"timestamp"`
	Added     []string      `json:"added"`
	Removed   []string      `json:"removed"`
	Modified  []string      `json:"modified"`
}

type PushCommitter struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
