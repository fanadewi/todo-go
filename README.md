# todo-go
Go-Learn

How to init a project

1. Create repo-name folder and initialize go mod and git
```
go mod init github.com/username/repo-name
git init
```
2. Create server.go as main file and write these default line

```
package main

import (
        "net/http"

        "github.com/gin-gonic/gin"
)

func main() {
        r := gin.Default()

        r.GET("/ping", func(c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{"reply": "pong"})
        })
        r.Run()
}
```

The basic ready to serve
