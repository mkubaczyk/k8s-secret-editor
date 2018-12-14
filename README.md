# Kubernetes Secret Editor

## Development

### Locally

#### Dependencies

`brew install dep` for golang package manager

Place this project in `$GOPATH/src/github.com/mkubaczyk/` directory

Run `dep ensure` to install packages in `vendor` directory

It still requires docker-compose to be built and run because of Redis and MySQL services dependency. Or do it on your own.

#### Run

`go run main.go`
