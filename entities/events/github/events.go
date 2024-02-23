package github

import "time"

type PushEventRequest struct {
	Source     string
	Type       string
	Ref        string `json:"ref,omitempty"`
	Before     string `json:"before,omitempty"`
	After      string `json:"after,omitempty"`
	Repository struct {
		ID       int    `json:"id,omitempty"`
		NodeID   string `json:"node_id,omitempty"`
		Name     string `json:"name,omitempty"`
		FullName string `json:"full_name,omitempty"`
		Private  bool   `json:"private,omitempty"`
		Owner    struct {
			Name              string `json:"name,omitempty"`
			Email             string `json:"email,omitempty"`
			Login             string `json:"login,omitempty"`
			ID                int    `json:"id,omitempty"`
			NodeID            string `json:"node_id,omitempty"`
			AvatarURL         string `json:"avatar_url,omitempty"`
			GravatarID        string `json:"gravatar_id,omitempty"`
			URL               string `json:"url,omitempty"`
			HTMLURL           string `json:"html_url,omitempty"`
			FollowersURL      string `json:"followers_url,omitempty"`
			FollowingURL      string `json:"following_url,omitempty"`
			GistsURL          string `json:"gists_url,omitempty"`
			StarredURL        string `json:"starred_url,omitempty"`
			SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
			OrganizationsURL  string `json:"organizations_url,omitempty"`
			ReposURL          string `json:"repos_url,omitempty"`
			EventsURL         string `json:"events_url,omitempty"`
			ReceivedEventsURL string `json:"received_events_url,omitempty"`
			Type              string `json:"type,omitempty"`
			SiteAdmin         bool   `json:"site_admin,omitempty"`
		} `json:"owner,omitempty"`
		HTMLURL                  string    `json:"html_url,omitempty"`
		Description              string    `json:"description,omitempty"`
		Fork                     bool      `json:"fork,omitempty"`
		URL                      string    `json:"url,omitempty"`
		ForksURL                 string    `json:"forks_url,omitempty"`
		KeysURL                  string    `json:"keys_url,omitempty"`
		CollaboratorsURL         string    `json:"collaborators_url,omitempty"`
		TeamsURL                 string    `json:"teams_url,omitempty"`
		HooksURL                 string    `json:"hooks_url,omitempty"`
		IssueEventsURL           string    `json:"issue_events_url,omitempty"`
		EventsURL                string    `json:"events_url,omitempty"`
		AssigneesURL             string    `json:"assignees_url,omitempty"`
		BranchesURL              string    `json:"branches_url,omitempty"`
		TagsURL                  string    `json:"tags_url,omitempty"`
		BlobsURL                 string    `json:"blobs_url,omitempty"`
		GitTagsURL               string    `json:"git_tags_url,omitempty"`
		GitRefsURL               string    `json:"git_refs_url,omitempty"`
		TreesURL                 string    `json:"trees_url,omitempty"`
		StatusesURL              string    `json:"statuses_url,omitempty"`
		LanguagesURL             string    `json:"languages_url,omitempty"`
		StargazersURL            string    `json:"stargazers_url,omitempty"`
		ContributorsURL          string    `json:"contributors_url,omitempty"`
		SubscribersURL           string    `json:"subscribers_url,omitempty"`
		SubscriptionURL          string    `json:"subscription_url,omitempty"`
		CommitsURL               string    `json:"commits_url,omitempty"`
		GitCommitsURL            string    `json:"git_commits_url,omitempty"`
		CommentsURL              string    `json:"comments_url,omitempty"`
		IssueCommentURL          string    `json:"issue_comment_url,omitempty"`
		ContentsURL              string    `json:"contents_url,omitempty"`
		CompareURL               string    `json:"compare_url,omitempty"`
		MergesURL                string    `json:"merges_url,omitempty"`
		ArchiveURL               string    `json:"archive_url,omitempty"`
		DownloadsURL             string    `json:"downloads_url,omitempty"`
		IssuesURL                string    `json:"issues_url,omitempty"`
		PullsURL                 string    `json:"pulls_url,omitempty"`
		MilestonesURL            string    `json:"milestones_url,omitempty"`
		NotificationsURL         string    `json:"notifications_url,omitempty"`
		LabelsURL                string    `json:"labels_url,omitempty"`
		ReleasesURL              string    `json:"releases_url,omitempty"`
		DeploymentsURL           string    `json:"deployments_url,omitempty"`
		CreatedAt                int       `json:"created_at,omitempty"`
		UpdatedAt                time.Time `json:"updated_at,omitempty"`
		PushedAt                 int       `json:"pushed_at,omitempty"`
		GitURL                   string    `json:"git_url,omitempty"`
		SSHURL                   string    `json:"ssh_url,omitempty"`
		CloneURL                 string    `json:"clone_url,omitempty"`
		SvnURL                   string    `json:"svn_url,omitempty"`
		Homepage                 string    `json:"homepage,omitempty"`
		Size                     int       `json:"size,omitempty"`
		StargazersCount          int       `json:"stargazers_count,omitempty"`
		WatchersCount            int       `json:"watchers_count,omitempty"`
		Language                 string    `json:"language,omitempty"`
		HasIssues                bool      `json:"has_issues,omitempty"`
		HasProjects              bool      `json:"has_projects,omitempty"`
		HasDownloads             bool      `json:"has_downloads,omitempty"`
		HasWiki                  bool      `json:"has_wiki,omitempty"`
		HasPages                 bool      `json:"has_pages,omitempty"`
		HasDiscussions           bool      `json:"has_discussions,omitempty"`
		ForksCount               int       `json:"forks_count,omitempty"`
		MirrorURL                any       `json:"mirror_url,omitempty"`
		Archived                 bool      `json:"archived,omitempty"`
		Disabled                 bool      `json:"disabled,omitempty"`
		OpenIssuesCount          int       `json:"open_issues_count,omitempty"`
		License                  any       `json:"license,omitempty"`
		AllowForking             bool      `json:"allow_forking,omitempty"`
		IsTemplate               bool      `json:"is_template,omitempty"`
		WebCommitSignoffRequired bool      `json:"web_commit_signoff_required,omitempty"`
		Topics                   []any     `json:"topics,omitempty"`
		Visibility               string    `json:"visibility,omitempty"`
		Forks                    int       `json:"forks,omitempty"`
		OpenIssues               int       `json:"open_issues,omitempty"`
		Watchers                 int       `json:"watchers,omitempty"`
		DefaultBranch            string    `json:"default_branch,omitempty"`
		Stargazers               int       `json:"stargazers,omitempty"`
		MasterBranch             string    `json:"master_branch,omitempty"`
	} `json:"repository,omitempty"`
	Pusher struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"pusher,omitempty"`
	Sender struct {
		Login             string `json:"login,omitempty"`
		ID                int    `json:"id,omitempty"`
		NodeID            string `json:"node_id,omitempty"`
		AvatarURL         string `json:"avatar_url,omitempty"`
		GravatarID        string `json:"gravatar_id,omitempty"`
		URL               string `json:"url,omitempty"`
		HTMLURL           string `json:"html_url,omitempty"`
		FollowersURL      string `json:"followers_url,omitempty"`
		FollowingURL      string `json:"following_url,omitempty"`
		GistsURL          string `json:"gists_url,omitempty"`
		StarredURL        string `json:"starred_url,omitempty"`
		SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
		OrganizationsURL  string `json:"organizations_url,omitempty"`
		ReposURL          string `json:"repos_url,omitempty"`
		EventsURL         string `json:"events_url,omitempty"`
		ReceivedEventsURL string `json:"received_events_url,omitempty"`
		Type              string `json:"type,omitempty"`
		SiteAdmin         bool   `json:"site_admin,omitempty"`
	} `json:"sender,omitempty"`
	Created bool   `json:"created,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
	Forced  bool   `json:"forced,omitempty"`
	BaseRef any    `json:"base_ref,omitempty"`
	Compare string `json:"compare,omitempty"`
	Commits []struct {
		ID        string    `json:"id,omitempty"`
		TreeID    string    `json:"tree_id,omitempty"`
		Distinct  bool      `json:"distinct,omitempty"`
		Message   string    `json:"message,omitempty"`
		Timestamp time.Time `json:"timestamp,omitempty"`
		URL       string    `json:"url,omitempty"`
		Author    struct {
			Name     string `json:"name,omitempty"`
			Email    string `json:"email,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"author,omitempty"`
		Committer struct {
			Name     string `json:"name,omitempty"`
			Email    string `json:"email,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"committer,omitempty"`
		Added    []any    `json:"added,omitempty"`
		Removed  []any    `json:"removed,omitempty"`
		Modified []string `json:"modified,omitempty"`
	} `json:"commits,omitempty"`
	HeadCommit struct {
		ID        string    `json:"id,omitempty"`
		TreeID    string    `json:"tree_id,omitempty"`
		Distinct  bool      `json:"distinct,omitempty"`
		Message   string    `json:"message,omitempty"`
		Timestamp time.Time `json:"timestamp,omitempty"`
		URL       string    `json:"url,omitempty"`
		Author    struct {
			Name     string `json:"name,omitempty"`
			Email    string `json:"email,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"author,omitempty"`
		Committer struct {
			Name     string `json:"name,omitempty"`
			Email    string `json:"email,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"committer,omitempty"`
		Added    []any    `json:"added,omitempty"`
		Removed  []any    `json:"removed,omitempty"`
		Modified []string `json:"modified,omitempty"`
	} `json:"head_commit,omitempty"`
}
