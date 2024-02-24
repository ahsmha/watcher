package crawl

type InitService interface {
	FetchAllRepos(token string)
	CreateHookToRepo(owner string, repo string, done chan struct{})
}
