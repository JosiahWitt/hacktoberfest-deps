package searchfest

import (
	"context"
	"net/url"
	"path"
	"strings"

	"github.com/google/go-github/v32/github"
)

type Listable interface {
	ListAllTopics(ctx context.Context, owner, repo string) ([]string, *github.Response, error)
}

type Result struct {
	RepoURL string
	Topics  []string
}

func SearchForHacktoberfestRepos(listable Listable, repos []string) ([]Result, error) {
	results := []Result{}

	for _, repoPath := range repos {
		if !strings.HasPrefix(repoPath, "https://github.com/") {
			continue
		}

		repoURL, err := url.Parse(repoPath)
		if err != nil {
			return nil, err
		}

		owner, repo := path.Split(repoURL.Path)
		owner = owner[1 : len(owner)-1] // Remove surrounding slashes
		topics, _, err := listable.ListAllTopics(context.Background(), owner, repo)
		if err != nil {
			return nil, err
		}

		if !containsString(topics, "hacktoberfest") {
			continue
		}

		results = append(results, Result{
			RepoURL: repoPath,
			Topics:  topics,
		})
	}

	return results, nil
}

func containsString(arr []string, str string) bool {
	for _, v := range arr {
		if strings.Contains(v, str) {
			return true
		}
	}

	return false
}
