# Why?

This is clean sandbox package to remind how to build, test, package go projects.

# How?
// see more at https://github.com/golang-standards/project-layout (sample: https://github.com/grafana/loki )

```
$ mkdir bago0
$ cd bago0
$ go mod init github.com/balionis-arvydas/balionis-go/bago0
$ mkdir -p cmd/bago0
$ touch cmd/bago0/main.go
```

# Build

```
$ go build ./...
```

# Test

```
$ go test ./...
```

# Run

```
$ go run ./...
// or
$ ./bago0.exe
```

