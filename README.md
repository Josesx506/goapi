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