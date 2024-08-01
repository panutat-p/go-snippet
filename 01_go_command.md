# Go Commands

## Go ENV

Get
```sh
go env GOROOT
```

Set
```sh
go env -w GO111MODULE=on
```

https://goproxy.io/docs/GOPRIVATE-env.html

```sh
go env -w GOPRIVATE='private-repo.com/*'
```

Unset
```sh
go env -u GO111MODULE
```

## Go Module

https://go.dev/ref/mod

* Go 1.16+, module-aware mode is enabled by default when `GO111MODULE=on` or `GO111MODULE=`
* Go 1.17+, module graph pruning, the go command avoids loading the complete module graph until (and unless) it is needed
* Go 1.17+, the go command adds all indirect requirements to `go.mod`

```sh
go mod init project-name
```

Install a dependency without updating other dependencies
```sh
go get github.com/joho/godotenv
```

Upgrade to the latest version if the package is already a dependency
```sh
go get -u github.com/joho/godotenv
```

Remove any unnecessary dependencies and ensure the cleanliness of your module files
```sh
go mod tidy
```

Download the dependencies specified in `go.mod`
```sh
go mod download
```

```sh
go list -m all
```

## Go Test

```sh
go test ./...
```

```sh
go test -coverprofile=coverage.out ./...
```

```sh
go tool cover -html=coverage.out
```
