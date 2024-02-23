package github

type PushEventRequest struct {
	Source  string
	Type    string
	Payload interface{}
}
