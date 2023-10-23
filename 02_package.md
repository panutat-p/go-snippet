# Package

https://go.dev/doc/modules/layout

## Package names

https://go.dev/blog/package-names

* Avoid meaningless: `util`, `common`, `misc`, ...
* Avoid single entry point: `api`, `types`, `interfaces`, ...
* Avoid unnecessary package name collisions

## Import

relative\
https://pkg.go.dev/cmd/go#hdr-Relative_import_paths

remote\
https://pkg.go.dev/cmd/go#hdr-Remote_import_paths

## Server layout

* Server projects typically won’t have packages for export
* It’s recommended to keep the Go packages implementing the server’s logic in the `internal` directory

```
project-root-directory/
  go.mod
  internal/
    auth/
      ...
    metrics/
      ...
    model/
      ...
  cmd/
    api-server/
      main.go
    metrics-analyzer/
      main.go
    ...
```

## Libraries

https://pkg.go.dev/github.com/geektime007/mgmt/pkg

https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal
