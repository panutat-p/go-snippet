# Go Commands

## Go ENV

Get
```shell
go env GOROOT
```

Set
```shell
go env -w GO111MODULE=on
```

Unset
```shell
go env -u GO111MODULE
```

## Go Module

https://go.dev/ref/mod

* Go 1.16+, module-aware mode is enabled by default when `GO111MODULE=on` or `GO111MODULE=`
* Go 1.17+, module graph pruning, the go command avoids loading the complete module graph until (and unless) it is needed
* Go 1.17+, the go command adds all indirect requirements to `go.mod`

```shell
go mod init project-name
```

```shell
go get github.com/joho/godotenv
```

```shell
go mod tidy
```

```shell
go list -m all
```

## Go Test
```shell
go test ./...
```
