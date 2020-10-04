package depfinder

import (
	"os/exec"
	"strings"

	"golang.org/x/tools/go/vcs"
)

func FindDependencyRepos() ([]string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", `{{ if not .Main }}{{ .Path }}{{ end }}`, "all")
	rawDepList, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	depRepos := []string{}
	inArray := map[string]bool{}
	for _, path := range strings.Split(string(rawDepList), "\n") {
		if strings.TrimSpace(path) == "" {
			continue
		}

		repoRoot, err := vcs.RepoRootForImportPath(path, false)
		if err != nil {
			return nil, err
		}
		repo := repoRoot.Repo

		if _, ok := inArray[repo]; ok {
			continue
		}

		inArray[repo] = true
		depRepos = append(depRepos, repo)
	}

	return depRepos, nil
}
