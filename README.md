# go-deck-builder

Go Deck Builder is a project I created to accomplish two things, learn Golang and create a webserver that builds Magic
the Gathering decks.

## Table of Contents
- [Architecture](#architecture)
- [Local Development](#local-development)

## Architecture
The Go Deck Builder integrates with the [Scryfall API](https://scryfall.com/docs/api), and uses this tool to build a deck for a user based on their
specified parameters.

## Local Development
1. Install Go locally, you can visit the official [Go downloads](https://golang.org/dl/) page to do that.
2. Install golangci-lint: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
    - You can verify the install by running `golangci-lint --version`

## Pull Requests
1. Before submitting a pull request run `golangcli-run lint`. Lint checks are run on every pull request and will fail if
there are any errors.