package controller

import (
	"net/http"
	"watcher/domain/event"
	"watcher/entities/events/github"
	"watcher/utils"

	"github.com/gin-gonic/gin"
)

type EventController interface {
	HandleEvent(ctx *gin.Context)
}

type eventControllerImplementation struct {
	service event.Service
}

func NewEventController(service event.Service) EventController {
	return &eventControllerImplementation{
		service: service,
	}
}

func (eci *eventControllerImplementation) HandleEvent(ctx *gin.Context) {

	// check if webhook repo has push webhook active
	if (isHookActive()) {

	} 
	// if not add it
	
	// get events

	var event interface{}
	var eventType string

	switch ctx.Request.Header.Get("User-Agent") {
		case utils.
	default:
		// log this event with no header
	}

	case utils.GITHUB_PUSH_EVENT:
		event = github.PushEventRequest{}
		break
	case utils.GITLAB_PUSH_EVENT:
		// gitlab
	if err := ctx.ShouldBindJSON(&event); err != nil {
		// add logging
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	xyz := eci.service.Push()
	mr.listener.Handle(c, event)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "event pushed to kafka successfully",
	})
}

func isHookActive() bool {
	return true
}

func (*service) initWatcher(?) ? {
	token:= getAccessToken()
	repos := GetAllUserRepos(token)
	return InitialProcessing(repos)
}

func (*rdb) getAccessToken(?) ? {
	// fetch from database
	return token
}

func GetAllUserRepos(?) ? {
	url := "https://api.github.com/repos/OWNER/REPO/hooks"
	method := "POST"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer <YOUR-TOKEN>")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	return ?
}
