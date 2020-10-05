# hacktoberfest-deps

Find your dependencies participating in [Hacktoberfest](https://hacktoberfest.digitalocean.com/).
Currently only supports Go dependencies.

## Installation

Download the [latest release](https://github.com/JosiahWitt/hacktoberfest-deps/releases).

## Usage

1. Navigate to your project root.
2. Run `hacktoberfest-deps`.
   - Set the `GITHUB_TOKEN` environment variable if you exceed the API limit.
     You can get a personal access token with `repo/public_repo` access by following [these directions](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token).

## Example

```bash
$ cd my/project # This directory has a go.mod file
$ GITHUB_TOKEN=my-token hacktoberfest-deps
Finding dependency repositories for current project...
Found 49 dependency repositories. Filtering by Hacktoberfest topic...

Filtered to 1 dependencies participating in Hacktoberfest:
 - https://github.com/JosiahWitt/erk [golang error-handling errors erk json-errors erk-errors hacktoberfest]
```
