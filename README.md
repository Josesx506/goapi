# goapi
First go api

- Install packages in your scripts that you don't have with `go mod tidy`
- Manually install package `go get github.com/gorilla/schema`
- Rename function on import in file
    ```go
    import (
        "errors"

        "github.com/github.com/Josesx506/goapi/api"
        log "github.com/sirupsen/logrus"
    )
    ```
- Start server with `go run cmd/api/main.go`
- Query the server from a different terminal
    ```bash
    # 200 response from mock db
    curl localhost:8000/account/coins/?username=jason \
        -H "Authorization: 456DEF"
    ```
    ```bash
    # 400 failed response
    curl localhost:8000/account/coins
    ```