package service

import (
	"fmt"
	"time"
	crawl "watcher/domain/init"
)

type Github struct {
	Token    string
	RepoChan chan (string)
}

func NewInitService(token string) crawl.InitService {
	return &Github{
		Token:    token,
		RepoChan: make(chan string),
	}
}
func (g *Github) FetchAllRepos(token string) {
	defer close(g.RepoChan)
	time.Sleep(time.Second * 2)
	//fetch all repos
	repos := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, repo := range repos {
		fmt.Println(repo)
		g.RepoChan <- repo
	}
}

func (g *Github) CreateHookToRepo(owner string, repo string, done chan struct{}) {
	//repoChan := make(chan string)
free:
	for {
		select {
		case rep, ok := <-g.RepoChan:
			if ok {
				//creating hook takes time
				time.Sleep(time.Second * 1)
				fmt.Println("Creating hook to repo", rep)
			} else {
				done <- struct{}{}
				break free
			}
		}
	}

}

func DoWork(serv crawl.InitService, done chan struct{}) {
	//serv := NewInitService("abc")
	go serv.FetchAllRepos("owner")
	go serv.CreateHookToRepo("owner", "repo", done)
}

// func main() {
// 	fmt.Println("runing service.init.go")
// 	serv := NewInitService("abc")
// 	go DoWork(serv)
// 	<-done
// 	//time.Sleep(time.Second * 20)

// }
