package main

import (
	"context"
	"fmt"
	"os"

	"github.com/JosiahWitt/hacktoberfest-deps/depfinder"
	"github.com/JosiahWitt/hacktoberfest-deps/searchfest"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func main() {
	gh := createGitHubClient()

	fmt.Println("Finding dependency repositories for current project...")
	repos, err := depfinder.FindDependencyRepos()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %v dependency repositories. Filtering by Hacktoberfest topic...\n", len(repos))
	results, err := searchfest.SearchForHacktoberfestRepos(gh.Repositories, repos)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("Filtered to %v dependencies participating in Hacktoberfest:\n", len(results))
	for _, result := range results {
		fmt.Printf(" - %s %+v\n", result.RepoURL, result.Topics)
	}
}

func createGitHubClient() *github.Client {
	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: ghToken},
		)
		tc := oauth2.NewClient(context.Background(), ts)

		return github.NewClient(tc)
	}

	fmt.Println(
		"WARNING: No GitHub access token set. " +
			"If you reach the API limit, please set an access token using the GITHUB_TOKEN environment variable. " +
			"Please provide repo/public_repo access. " +
			"See: https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token",
	)
	fmt.Println()
	return github.NewClient(nil)
}
