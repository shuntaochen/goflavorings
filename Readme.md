## GO Flavorings,

1. go mod init' as 'goflavorings'
2. in main, appName(), version(), 

Prerequsites:

```
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

Hints and musts:

```
1. fileName having _, will not be included in project files, such as 'session_test.go, this has to be run by go run session_test.go
2. reflex: reflex -r "\.go$" -s go run .
```

TODOs:
1. go session, get package by: go get github.com/kataras/go-sessions/v3
2. go cache
3. go jwt
4. go json
5. routes.go
6. go reflex, 'go get github.com/cespare/reflex
