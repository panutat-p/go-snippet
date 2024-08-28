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

Unset
```sh
go env -u GO111MODULE
```

Exclude the package from the checksum database\
```sh
go env -w GOPRIVATE='private-repo.com/*'
```

Specify the proxy server that serving the private modules
```sh
go env -w GOPROXY='proxy.example.com'
```

## Go Toolchains

https://tip.golang.org/doc/toolchain

Update the module to require the latest released Go toolchain
```sh
go get go@latest
```

## Go Module

https://go.dev/ref/mod

https://goproxy.io/docs/GOPRIVATE-env.html

* Go 1.16+, module-aware mode is enabled by default when `GO111MODULE=on` or `GO111MODULE=`
* Go 1.17+, module graph pruning, the go command avoids loading the complete module graph until (and unless) it is needed
* Go 1.17+, the go command adds all indirect requirements to `go.mod`

```sh
go mod init project-name
```

```sh
go mod edit -go 1.23.0
```

Install or update a dependency without updating other dependencies
```sh
go get github.com/labstack/echo/v4
```

Upgrade to the latest version including its dependencies
```sh
go get -u github.com/labstack/echo/v4
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

Shows a shortest path in the import graph
```sh
go mod why github.com/labstack/echo/v4
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
