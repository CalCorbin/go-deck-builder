# go-deck-builder

Go Deck Builder is a project I created to accomplish two things, learn Golang and create a webserver that builds Magic the Gathering decks.

## Table of Contents
* [go-deck-builder](#go-deck-builder)
  * [Table of Contents](#table-of-contents)
  * [Architecture](#architecture)
  * [Local Development](#local-development)
    * [Minimum Requirements](#minimum-requirements)
    * [Recommended Requirements](#recommended-requirements)
  * [Pull Requests](#pull-requests)

## Architecture

The Go Deck Builder integrates with the [Scryfall API](https://scryfall.com/docs/api), and uses this tool to build a deck for a user based on their specified parameters.

## Local Development

### Minimum Requirements

1. Install Go locally, you can visit the official [Go downloads](https://golang.org/dl/) page to do that.
2. Install golangci-lint: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
    - You can verify the install by running `golangci-lint --version`

### Recommended Requirements

* Install [pre-commit hooks](https://pre-commit.com/) to run lint checks before every commit.
  * Install pre-commit: `brew install pre-commit`
  * Install the pre-commit hooks: `pre-commit install`

## Pull Requests

1. Before submitting a pull request run `golangci-lint run`. Lint checks are run on every pull request and will fail if there are any errors.
