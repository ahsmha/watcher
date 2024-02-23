package main

import (
	"watcher/api"
	"watcher/service"
	"watcher/utils"
)

func main() {
	// logger := utils.NewLogger(config)
	config := utils.NewConfig()
	services := service.InitServices(config /*logger,*/)
	api.StartServer(services, config)
}

// func InitialProcessing([]string repos) bool {
// 	// processing the commit data for each repos.
// 	GetAllRepoCommits(?)
// 	// for each commit data, sum additions
// 	// save totalcommit sum in db
// 	return true
// }

// func GetAllRepoCommits(?) ? {
// }
