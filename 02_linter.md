# Linters

### gofmt

https://pkg.go.dev/cmd/gofmt

```sh
gofmt -w .
```

## goimports

https://pkg.go.dev/golang.org/x/tools/cmd/goimports

```sh
go install golang.org/x/tools/cmd/goimports@latest
```

```sh
goimports -w .
```

## gci

https://github.com/daixiang0/gci

```sh
go install github.com/daixiang0/gci@latest
```

```sh
gci write .
```

```sh
gci write -s standard -s default -s 'prefix(github.com/panutat-p)' -s localmodule .
```

## golangci-lint

```sh
golangci-lint run ./...
```

```sh
golangci-lint run --fix ./...
```

https://github.com/golangci/golangci-lint

```yaml
run:
  concurrency: 4
  timeout: 5m

output:
  formats:
    - format: colored-line-number
      path: stderr
    - format: html
      path: golangci_lint_report.html

linters:
  enable:
    - errcheck
    - gocyclo
    - gosec
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - unused
    - gci

linters-settings:
  gci:
    sections:
      - standard
      - default
      - prefix(github.com/panutat-p)
      - localmodule
    skip-generated: true
    custom-order: true

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
  exclude-rules:
    - linters:
        - gosec
      text: "G115: integer overflow conversion*"
    - linters:
        - errcheck
      text: "Error return value of .*Close.* is not checked"
```
