package main

import (
	"watcher/api"
	"watcher/service"
)

func main() {
	// logger := utils.NewLogger(config)
	// check if webhook repo has push webhook active
	// if not add it
	// get events
	services := service.InitServices( /*logger,*/ )
	api.StartServer(services)
}

// func (*service) initWatcher(?) ? {
// 	token:= getAccessToken()
// 	repos := GetAllUserRepos(token)
// 	return InitialProcessing(repos)
// }

// func (*rdb) getAccessToken(?) ? {
// 	// fetch from database
// 	return token
// }

// func GetAllUserRepos(?) ? {
// 	url := "https://api.github.com/repos/OWNER/REPO/hooks"
// 	method := "POST"

// 	payload := strings.NewReader("")

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	req.Header.Add("Accept", "application/vnd.github+json")
// 	req.Header.Add("Authorization", "Bearer <YOUR-TOKEN>")
// 	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(body))

// 	return ?
// }

// func InitialProcessing([]string repos) bool {
// 	// processing the commit data for each repos.
// 	GetAllRepoCommits(?)
// 	// for each commit data, sum additions
// 	// save totalcommit sum in db
// 	return true
// }

// func GetAllRepoCommits(?) ? {
// }
