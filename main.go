package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
			<!doctype html>
			<html lang="ru">
			<head>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1" />
			<title>Привет</title>
			</head>
			<body style="font-family: system-ui; padding: 40px;">
			<h1>Hello GO</h1>
			
			</body>
			</html>
		`))
	})

	r.Run(":8080")
}
